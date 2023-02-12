package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/middlewares"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
)

// BindModelReadmeApi registers the admin api endpoints and the corresponding handlers.
func BindModelReadmeApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	modelGroup := rg.Group("/org/:orgId/model", middlewares.RequireAuthContext, middlewares.ValidateOrg(api.app))
	modelGroup.GET("/:modelName/readme/version/:version", api.DefaultHandler(GetModelReadmeVersion), middlewares.ValidateModel(api.app))
	modelGroup.GET("/:modelName/readme/version", api.DefaultHandler(GetModelReadmeAllVersions), middlewares.ValidateModel(api.app))
	modelGroup.POST("/:modelName/readme", api.DefaultHandler(UpdateModelReadme), middlewares.ValidateModel(api.app))
}

// GetModelReadmeAllVersions godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get model readme
//	@Description	Get model readme
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/readme/version [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
func (api *Api) GetModelReadmeAllVersions(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	readme, err := api.app.Dao().GetModelReadmeAllVersions(modelUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	response := models.NewDataResponse(http.StatusOK, readme, "Model Readme version")
	return response
}

// GetModelReadmeVersion godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get model readme
//	@Description	Get model readme
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/readme/version/{version} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
//	@Param			version		path	string	true	"Version"
func (api *Api) GetModelReadmeVersion(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	versionName := request.GetPathParam("version")
	readme, err := api.app.Dao().GetModelReadmeVersion(modelUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if readme == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Model readme version not found")
	}
	response := models.NewDataResponse(http.StatusOK, readme, "Model Readme version")
	return response
}

// UpdateModelReadme godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update readme of a model for a category
//	@Description	Update readme of a model for a category
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/readme [post]
//	@Param			orgId		path	string					true	"Organization Id"
//	@Param			modelName	path	string					true	"Model Name"
//	@Param			data		body	models.ReadmeRequest	true	"Data"
func (api *Api) UpdateModelReadme(request *models.Request) *models.Response {
	request.ParseJsonBody()
	modelUUID := request.GetModelUUID()
	modelFileType := request.GetParsedBodyAttribute("file_type")
	var modelFileTypeData string
	if modelFileType == nil {
		modelFileTypeData = ""
	} else {
		modelFileTypeData = modelFileType.(string)
	}
	modelContent := request.GetParsedBodyAttribute("content")
	var modelContentData string
	if modelContent == nil {
		modelContentData = ""
	} else {
		modelContentData = modelContent.(string)
	}
	readme, err := api.app.Dao().UpdateModelReadme(modelUUID, modelFileTypeData, modelContentData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, readme, "Model readme updated")
}

var GetModelReadmeAllVersions ServiceFunc = (*Api).GetModelReadmeAllVersions
var GetModelReadmeVersion ServiceFunc = (*Api).GetModelReadmeVersion
var UpdateModelReadme ServiceFunc = (*Api).UpdateModelReadme
