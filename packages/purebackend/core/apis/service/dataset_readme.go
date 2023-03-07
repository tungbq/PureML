package service

import (
	"net/http"

	authmiddlewares "github.com/PuremlHQ/PureML/packages/purebackend/auth/middlewares"
	"github.com/PuremlHQ/PureML/packages/purebackend/core"
	"github.com/PuremlHQ/PureML/packages/purebackend/core/models"
	"github.com/PuremlHQ/PureML/packages/purebackend/dataset/middlewares"
	orgmiddlewares "github.com/PuremlHQ/PureML/packages/purebackend/org/middlewares"
	"github.com/labstack/echo/v4"
)

// BindDatasetReadmeApi registers the admin api endpoints and the corresponding handlers.
func BindDatasetReadmeApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	datasetGroup := rg.Group("/org/:orgId/dataset", authmiddlewares.RequireAuthContext, orgmiddlewares.ValidateOrg(api.app))
	datasetGroup.GET("/:datasetName/readme/version/:version", api.DefaultHandler(GetDatasetReadmeVersion), middlewares.ValidateDataset(api.app))
	datasetGroup.GET("/:datasetName/readme/version", api.DefaultHandler(GetDatasetReadmeAllVersions), middlewares.ValidateDataset(api.app))
	datasetGroup.POST("/:datasetName/readme", api.DefaultHandler(UpdateDatasetReadme), middlewares.ValidateDataset(api.app))
}

// GetDatasetReadmeAllVersions godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get dataset readme
//	@Description	Get dataset readme
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/readme/version [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
func (api *Api) GetDatasetReadmeAllVersions(request *models.Request) *models.Response {
	modelUUID := request.GetDatasetUUID()
	readme, err := api.app.Dao().GetDatasetReadmeAllVersions(modelUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	response := models.NewDataResponse(http.StatusOK, readme, "Dataset Readme version")
	return response
}

// GetDatasetReadmeVersion godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get dataset readme
//	@Description	Get dataset readme
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/readme/version/{version} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			version		path	string	true	"Version"
func (api *Api) GetDatasetReadmeVersion(request *models.Request) *models.Response {
	modelUUID := request.GetDatasetUUID()
	versionName := request.GetPathParam("version")
	readme, err := api.app.Dao().GetDatasetReadmeVersion(modelUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if readme == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Dataset readme version not found")
	}
	response := models.NewDataResponse(http.StatusOK, readme, "Dataset Readme version")
	return response
}

// UpdateDatasetReadme godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update readme of a dataset for a category
//	@Description	Update readme of a dataset for a category
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/readme [post]
//	@Param			orgId		path	string					true	"Organization Id"
//	@Param			datasetName	path	string					true	"Dataset Name"
//	@Param			data		body	models.ReadmeRequest	true	"Data"
func (api *Api) UpdateDatasetReadme(request *models.Request) *models.Response {
	request.ParseJsonBody()
	datasetUUID := request.GetDatasetUUID()
	datasetFileType := request.GetParsedBodyAttribute("file_type")
	var datasetFileTypeData string
	if datasetFileType == nil {
		datasetFileTypeData = ""
	} else {
		datasetFileTypeData = datasetFileType.(string)
	}
	datasetContent := request.GetParsedBodyAttribute("content")
	var datasetContentData string
	if datasetContent == nil {
		datasetContentData = ""
	} else {
		datasetContentData = datasetContent.(string)
	}
	readme, err := api.app.Dao().UpdateDatasetReadme(datasetUUID, datasetFileTypeData, datasetContentData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, readme, "Dataset readme updated")
}

var GetDatasetReadmeAllVersions ServiceFunc = (*Api).GetDatasetReadmeAllVersions
var GetDatasetReadmeVersion ServiceFunc = (*Api).GetDatasetReadmeVersion
var UpdateDatasetReadme ServiceFunc = (*Api).UpdateDatasetReadme
