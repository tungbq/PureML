package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func CreateModelActivity(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	modelName := request.GetPathParam("modelName")
	activityName := request.GetPathParam("activityName")
	activity, err := datastore.GetModelActivity(orgId, modelName, activityName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if activity != nil {
		return models.NewDataResponse(http.StatusConflict, nil, "Activity with same name already exists")
	}
	createdActivity, err := datastore.CreateModelActivity(orgId, modelName, activityName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*createdActivity}, "Activity created")
}
