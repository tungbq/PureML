package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetModelActivity(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	modelName := request.GetPathParam("modelName")
	activityName := request.GetPathParam("activityName")
	activity, err := datastore.GetModelActivity(orgId, modelName, activityName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if activity == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Activity not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*activity}, "Activity found")
}
