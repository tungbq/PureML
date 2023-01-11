package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetActivity(request *models.Request) *models.Response {
	modelName := request.GetPathParam("modelName")
	datasetName := request.GetPathParam("datasetName")
	activityName := request.GetPathParam("activityName")
	activity, err := datastore.GetActivity(activityName, modelName, datasetName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if activity == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Activity not found")
	}
	return models.NewDataResponse(http.StatusOK, activity, "Activity found")
}
