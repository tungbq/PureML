package service

import (
	_ "fmt"
	"net/http"
	"strings"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// RegisterModel godoc
// @Security ApiKeyAuth
// @Summary Register model
// @Description Register model file. Create model and default branches if not exists
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/model/{modelName}/branch/{branchName}/register [post]
// @Param file formData file true "Model file"
// @Param orgId path string true "Organization UUID"
// @Param modelName path string true "Model name"
// @Param branchName path string true "Branch name"
// @Param data formData models.RegisterModelRequest true "Model details"
func RegisterModel(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	var modelHash string
	if request.FormValues["hash"] != nil && len(request.FormValues["hash"]) > 0 {
		modelHash = request.FormValues["hash"][0]
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Hash is required")
	}
	var modelSourceType string
	if request.FormValues["storage"] != nil && len(request.FormValues["storage"]) > 0 {
		modelSourceType = strings.ToUpper(request.FormValues["storage"][0])
	}
	var modelIsEmpty bool
	if request.FormValues["isEmpty"] != nil && len(request.FormValues["isEmpty"]) > 0 {
		modelIsEmpty = request.FormValues["isEmpty"][0] == "true"
	}
	fileHeader := request.GetFormFile("file")
	if fileHeader == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "File is required")
	}
	modelBranchName := request.GetPathParam("branchName")
	if modelBranchName == "main" {
		return models.NewErrorResponse(http.StatusBadRequest, "Cannot register model directly to main branch")
	}
	sourceValid := false
	for source := range models.SupportedSources {
		if models.SupportedSources[source] == modelSourceType {
			sourceValid = true
			break
		}
	}
	if !sourceValid {
		return models.NewErrorResponse(http.StatusBadRequest, "Unsupported model source type")
	}
	modelBranchUUID := request.GetModelBranchUUID()
	modelVersion, err := datastore.UploadAndRegisterModelFile(orgId, modelBranchUUID, fileHeader, modelIsEmpty, modelHash, modelSourceType)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, modelVersion, "Model successfully registered")
}
