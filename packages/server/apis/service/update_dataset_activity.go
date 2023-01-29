package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// UpdateDatasetActivity godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update activity of a dataset for a category
//	@Description	Update activity of a dataset for a category
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/activity/{category}/{activityUUID} [post]
//	@Param			orgId			path	string	true	"Organization Id"
//	@Param			datasetName		path	string	true	"Dataset Name"
//	@Param			category		path	string	true	"Category"
//	@Param			activityUUID	path	string	true	"Activity UUID"
//	@Param			activity		body	string	true	"Activity"
func UpdateDatasetActivity(request *models.Request) *models.Response {
	request.ParseJsonBody()
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	updatedActivity := request.GetParsedBodyAttribute("activity").(string)
	updatedAttributes := map[string]string{}
	if updatedActivity != "" {
		updatedAttributes["activity"] = updatedActivity
	}
	updatedDbActivity, err := datastore.UpdateDatasetActivity(activityUUID, updatedAttributes)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*updatedDbActivity}, "Dataset Activity updated")
}
