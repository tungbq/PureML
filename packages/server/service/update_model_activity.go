package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// UpdateModelActivity godoc
// @Security ApiKeyAuth
// @Summary Update activity of a model for a category
// @Description Update activity of a model for a category
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/activity/{category}/{activityUUID} [get]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param category path string true "Category"
// @Param activityUUID path string true "Activity UUID"
func UpdateModelActivity(request *models.Request) *models.Response {
	request.ParseJsonBody()
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	updatedActivity := request.GetParsedBodyAttribute("activity").(string)
	updatedAttributes := map[string]string{}
	if updatedActivity != "" {
		updatedAttributes["activity"] = updatedActivity
	}
	updatedDbActivity, err := datastore.UpdateModelActivity(activityUUID, updatedAttributes)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*updatedDbActivity}, "Model Activity updated")
}
