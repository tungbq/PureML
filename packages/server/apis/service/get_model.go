package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetModel godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get specific model of an organization
//	@Description	Get specific model of an organization
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
func GetModel(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	modelName := request.GetModelName()
	model, err := datastore.GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Model not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.ModelResponse{*model}, "Model successfully retrieved")
}
