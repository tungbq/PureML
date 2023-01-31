package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// GetModelActivity godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get activity of a model for a category
//	@Description	Get activity of a model for a category
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/activity/{category} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
//	@Param			category	path	string	true	"Category"
func GetModelActivity(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	category := request.GetPathParam("category")
	activity, err := datastore.GetModelActivity(modelUUID, category)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if activity == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Activity not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*activity}, "Activity found")
}

// CreateModelActivity godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Add activity of a model for a category
//	@Description	Add activity of a model for a category
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/activity/{category} [post]
//	@Param			orgId		path	string					true	"Organization Id"
//	@Param			modelName	path	string					true	"Model Name"
//	@Param			category	path	string					true	"Category"
//	@Param			data		body	models.ActivityRequest	true	"Activity"
func CreateModelActivity(request *models.Request) *models.Response {
	request.ParseJsonBody()
	modelUUID := request.GetModelUUID()
	userUUID := request.GetUserUUID()
	category := request.GetPathParam("category")
	activity := request.GetParsedBodyAttribute("activity")
	if activity == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Activity not found in request body")
	} else if activity.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Activity cannot be empty")
	}
	activityData := activity.(string)
	createdActivity, err := datastore.CreateModelActivity(modelUUID, userUUID, category, activityData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*createdActivity}, "Activity created")
}

// UpdateModelActivity godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update activity of a model for a category
//	@Description	Update activity of a model for a category
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/activity/{category}/{activityUUID} [post]
//	@Param			orgId			path	string	true	"Organization Id"
//	@Param			modelName		path	string	true	"Model Name"
//	@Param			category		path	string	true	"Category"
//	@Param			activityUUID	path	string	true	"Activity UUID"
//	@Param			activity		body	string	true	"Activity"
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

// DeleteModelActivity godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete an activity of a model for a category
//	@Description	Delete an activity of a model for a category
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/activity/{category}/{activityUUID}/delete [delete]
//	@Param			orgId			path	string	true	"Organization Id"
//	@Param			modelName		path	string	true	"Model Name"
//	@Param			category		path	string	true	"Category"
//	@Param			activityUUID	path	string	true	"Activity UUID"
func DeleteModelActivity(request *models.Request) *models.Response {
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	err := datastore.DeleteModelActivity(activityUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Model Activity deleted")
}
