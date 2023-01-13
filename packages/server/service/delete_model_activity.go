package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func DeleteModelActivity(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	modelName := request.GetPathParam("modelName")
	activityName := request.GetPathParam("activityName")
	_, err := datastore.DeleteModelActivity(orgId, modelName, activityName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Model Activity deleted")
}
