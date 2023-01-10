package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetModel(request *models.Request) *models.Response {
	orgId := request.GetPathParam("orgId")
	modelName := request.GetPathParam("modelName")
	model, err := datastore.GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Model not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.ModelNameResponse{*model}, "Model successfully retrieved")
}
