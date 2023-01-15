package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// DeleteModelActivity godoc
// @Security ApiKeyAuth
// @Summary Delete an activity of a model for a category
// @Description Delete an activity of a model for a category
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/activity/{category}/{activityUUID}/delete [delete]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param category path string true "Category"
// @Param activityUUID path string true "Activity UUID"
func DeleteModelActivity(request *models.Request) *models.Response {
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	err := datastore.DeleteModelActivity(activityUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Model Activity deleted")
}
