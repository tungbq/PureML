package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/middlewares"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// BindModelActivityApi registers the admin api endpoints and the corresponding handlers.
func BindModelActivityApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	modelGroup := rg.Group("/org/:orgId/model", middlewares.AuthenticateJWT(api.app), middlewares.ValidateOrg(api.app))
	modelGroup.GET("/:modelName/activity/:category", api.DefaultHandler(GetModelActivity), middlewares.ValidateModel(api.app))
	modelGroup.POST("/:modelName/activity/:category", api.DefaultHandler(CreateModelActivity), middlewares.ValidateModel(api.app))
	modelGroup.POST("/:modelName/activity/:category/:activityUUID", api.DefaultHandler(UpdateModelActivity), middlewares.ValidateModel(api.app))
	modelGroup.DELETE("/:modelName/activity/:category/:activityUUID/delete", api.DefaultHandler(DeleteModelActivity), middlewares.ValidateModel(api.app))
}

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
func (api *Api) GetModelActivity(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	category := request.GetPathParam("category")
	activity, err := api.app.Dao().GetModelActivity(modelUUID, category)
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
func (api *Api) CreateModelActivity(request *models.Request) *models.Response {
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
	createdActivity, err := api.app.Dao().CreateModelActivity(modelUUID, userUUID, category, activityData)
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
func (api *Api) UpdateModelActivity(request *models.Request) *models.Response {
	request.ParseJsonBody()
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	updatedActivity := request.GetParsedBodyAttribute("activity").(string)
	updatedAttributes := map[string]string{}
	if updatedActivity != "" {
		updatedAttributes["activity"] = updatedActivity
	}
	updatedDbActivity, err := api.app.Dao().UpdateModelActivity(activityUUID, updatedAttributes)
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
func (api *Api) DeleteModelActivity(request *models.Request) *models.Response {
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	err := api.app.Dao().DeleteModelActivity(activityUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Model Activity deleted")
}

var GetModelActivity ServiceFunc = (*Api).GetModelActivity
var CreateModelActivity ServiceFunc = (*Api).CreateModelActivity
var UpdateModelActivity ServiceFunc = (*Api).UpdateModelActivity
var DeleteModelActivity ServiceFunc = (*Api).DeleteModelActivity
