package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// GetDatasetActivity godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get activity of a dataset for a category
//	@Description	Get activity of a dataset for a category
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/activity/{category} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			category	path	string	true	"Category"
func GetDatasetActivity(request *models.Request) *models.Response {
	datasetUUID := request.GetDatasetUUID()
	category := request.GetPathParam("category")
	activity, err := datastore.GetDatasetActivity(datasetUUID, category)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if activity == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Activity not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*activity}, "Activity found")
}

// CreateDatasetActivity godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Add activity of a dataset for a category
//	@Description	Add activity of a dataset for a category
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/activity/{category} [post]
//	@Param			orgId		path	string					true	"Organization Id"
//	@Param			datasetName	path	string					true	"Dataset Name"
//	@Param			category	path	string					true	"Category"
//	@Param			data		body	models.ActivityRequest	true	"Activity"
func CreateDatasetActivity(request *models.Request) *models.Response {
	request.ParseJsonBody()
	datasetUUID := request.GetDatasetUUID()
	userUUID := request.GetUserUUID()
	category := request.GetPathParam("category")
	activity := request.GetParsedBodyAttribute("activity")
	if activity == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Activity not found in request body")
	} else if activity.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Activity cannot be empty")
	}
	activityData := activity.(string)
	createdActivity, err := datastore.CreateDatasetActivity(datasetUUID, userUUID, category, activityData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*createdActivity}, "Activity created")
}

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

// DeleteDatasetActivity godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete an activity of a dataset for a category
//	@Description	Delete an activity of a dataset for a category
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/activity/{category}/{activityUUID}/delete [delete]
//	@Param			orgId			path	string	true	"Organization Id"
//	@Param			datasetName		path	string	true	"Dataset Name"
//	@Param			category		path	string	true	"Category"
//	@Param			activityUUID	path	string	true	"Activity UUID"
func DeleteDatasetActivity(request *models.Request) *models.Response {
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	err := datastore.DeleteDatasetActivity(activityUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Dataset Activity deleted")
}
