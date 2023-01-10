package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

func GetModel(request *models.Request) *models.Response {
	orgId := request.GetPathParam("orgId")
	modelName := request.GetPathParam("modelName")
	model, err := datastore.GetModelByName(uuid.Must(uuid.FromString(orgId)), modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Model not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.ModelResponse{*model}, "Model successfully retrieved")
}
