package service

import (
	_ "fmt"
	"net/http"
	"strings"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// RegisterDataset godoc
// @Security ApiKeyAuth
// @Summary Register dataset
// @Description Register dataset file. Create dataset and default branches if not exists
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/branch/{branchName}/register [post]
// @Param file formData file true "Dataset file"
// @Param orgId path string true "Organization UUID"
// @Param datasetName path string true "Dataset name"
// @Param branchName path string true "Branch name"
// @Param data formData models.RegisterDatasetRequest true "Dataset details"
func RegisterDataset(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	var datasetHash string
	if request.FormValues["hash"] != nil && len(request.FormValues["hash"]) > 0 {
		datasetHash = request.FormValues["hash"][0]
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Hash is required")
	}
	var datasetSourceType string
	if request.FormValues["storage"] != nil && len(request.FormValues["storage"]) > 0 {
		datasetSourceType = strings.ToUpper(request.FormValues["storage"][0])
	}
	var datasetIsEmpty bool
	if request.FormValues["is_empty"] != nil && len(request.FormValues["is_empty"]) > 0 {
		datasetIsEmpty = request.FormValues["is_empty"][0] == "true"
	}
	var datasetLineage string
	if request.FormValues["lineage"] != nil && len(request.FormValues["lineage"]) > 0 {
		datasetLineage = request.FormValues["lineage"][0]
	}
	fileHeader := request.GetFormFile("file")
	if fileHeader == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "File is required")
	}
	datasetBranchName := request.GetPathParam("branchName")
	if datasetBranchName == "main" {
		return models.NewErrorResponse(http.StatusBadRequest, "Cannot register model directly to main branch")
	}
	sourceValid := false
	for source := range models.SupportedSources {
		if models.SupportedSources[source] == datasetSourceType {
			sourceValid = true
			break
		}
	}
	if !sourceValid {
		return models.NewErrorResponse(http.StatusBadRequest, "Unsupported model source type")
	}
	datasetBranchUUID := request.GetDatasetBranchUUID()
	datasetVersion, err := datastore.UploadAndRegisterDatasetFile(orgId, datasetBranchUUID, fileHeader, datasetIsEmpty, datasetHash, datasetSourceType, datasetLineage)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, datasetVersion, "Dataset successfully registered")
}
