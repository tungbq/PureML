package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

var defaultBranchNames = []string{"main", "developement"}

// RegisterModel godoc
// @Security ApiKeyAuth
// @Summary Register model
// @Description Register model file. Create model and default branches if not exists
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/:orgId/model/:modelName/register [post]
// @Param orgId path string true "Organization UUID"
// @Param modelName path string true "Model name"
// @Param data body models.RegisterModelRequest true "Model details"
func RegisterModel(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetOrgId()
	userUUID := request.GetUserUUID()
	modelName := request.GetPathParam("modelName")
	modelWiki := request.GetParsedBodyAttribute("wiki").(string)
	modelHash := request.GetParsedBodyAttribute("hash").(string)
	model, err := datastore.GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model == nil {
		model, err := datastore.CreateModel(orgId, modelName, modelWiki, userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		modelBranches, err := datastore.CreateModelBranches(model.UUID, defaultBranchNames)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		fileHeader := request.GetFormFile("file")
		if fileHeader == nil {
			//TODO mandatory file is missing return with appropriate response here
			return models.NewErrorResponse(http.StatusBadRequest, "File is required")
		}
		modelVersion, err := datastore.UploadAndRegisterModelFile(modelBranches[1].UUID, fileHeader, modelHash, "R2")
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		return models.NewDataResponse(http.StatusOK, modelVersion, "Model successfully created")

	}
	return models.NewErrorResponse(http.StatusConflict, "Model with same name already exists")
}
