package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func UpdateModelActivity(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	modelName := request.GetPathParam("modelName")
	activityName := request.GetPathParam("activityName")
	request.ParseJsonBody()
	updatedActivityName := request.GetParsedBodyAttribute("activityName").(string)
	updatedAttributes := map[string]string{}
	if updatedActivityName != "" && updatedActivityName != activityName {
		updatedAttributes["activityName"] = updatedActivityName
	}
	updatedActivity, err := datastore.UpdateModelActivity(orgId, modelName, activityName, updatedAttributes)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*updatedActivity}, "Model Activity updated")
}
