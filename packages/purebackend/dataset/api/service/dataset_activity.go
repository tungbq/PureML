package service

import (
	"net/http"

	authmiddlewares "github.com/PureMLHQ/PureML/packages/purebackend/auth/middlewares"
	"github.com/PureMLHQ/PureML/packages/purebackend/core"
	"github.com/PureMLHQ/PureML/packages/purebackend/core/models"
	"github.com/PureMLHQ/PureML/packages/purebackend/dataset/middlewares"
	orgmiddlewares "github.com/PureMLHQ/PureML/packages/purebackend/user_org/middlewares"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// BindDatasetActivityApi registers the dataset activity api endpoints and the corresponding handlers.
func BindDatasetActivityApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	datasetGroup := rg.Group("/org/:orgId/dataset", authmiddlewares.RequireAuthContext, orgmiddlewares.ValidateOrg(api.app))
	datasetGroup.GET("/:datasetName/activity/:category", api.DefaultHandler(GetDatasetActivity), middlewares.ValidateDataset(api.app))
	datasetGroup.POST("/:datasetName/activity/:category", api.DefaultHandler(CreateDatasetActivity), middlewares.ValidateDataset(api.app))
	datasetGroup.POST("/:datasetName/activity/:category/:activityUUID", api.DefaultHandler(UpdateDatasetActivity), middlewares.ValidateDataset(api.app))
	datasetGroup.DELETE("/:datasetName/activity/:category/:activityUUID/delete", api.DefaultHandler(DeleteDatasetActivity), middlewares.ValidateDataset(api.app))
}

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
func (api *Api) GetDatasetActivity(request *models.Request) *models.Response {
	datasetUUID := request.GetDatasetUUID()
	category := request.GetPathParam("category")
	activity, err := api.app.Dao().GetDatasetActivity(datasetUUID, category)
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
func (api *Api) CreateDatasetActivity(request *models.Request) *models.Response {
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
	createdActivity, err := api.app.Dao().CreateDatasetActivity(datasetUUID, userUUID, category, activityData)
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
func (api *Api) UpdateDatasetActivity(request *models.Request) *models.Response {
	request.ParseJsonBody()
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	updatedActivity := request.GetParsedBodyAttribute("activity").(string)
	updatedAttributes := map[string]string{}
	if updatedActivity != "" {
		updatedAttributes["activity"] = updatedActivity
	}
	updatedDbActivity, err := api.app.Dao().UpdateDatasetActivity(activityUUID, updatedAttributes)
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
func (api *Api) DeleteDatasetActivity(request *models.Request) *models.Response {
	activityUUID := uuid.Must(uuid.FromString(request.GetPathParam("activityUUID")))
	err := api.app.Dao().DeleteDatasetActivity(activityUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Dataset Activity deleted")
}

var GetDatasetActivity ServiceFunc = (*Api).GetDatasetActivity
var CreateDatasetActivity ServiceFunc = (*Api).CreateDatasetActivity
var UpdateDatasetActivity ServiceFunc = (*Api).UpdateDatasetActivity
var DeleteDatasetActivity ServiceFunc = (*Api).DeleteDatasetActivity
