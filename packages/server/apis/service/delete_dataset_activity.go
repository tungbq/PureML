package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// DeleteDatasetActivity godoc
// @Security ApiKeyAuth
// @Summary Delete an activity of a dataset for a category
// @Description Delete an activity of a dataset for a category
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/dataset/{datasetName}/activity/{category}/{activityUUID}/delete [delete]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param category path string true "Category"
// @Param activityUUID path string true "Activity UUID"
func DeleteDatasetActivity(request *models.Request) *models.Response {
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	err := datastore.DeleteDatasetActivity(activityUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Dataset Activity deleted")
}
