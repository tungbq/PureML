package service

import (
	"net/http"

	authmiddlewares "github.com/PureMLHQ/PureML/packages/purebackend/auth/middlewares"
	"github.com/PureMLHQ/PureML/packages/purebackend/core"
	"github.com/PureMLHQ/PureML/packages/purebackend/core/models"
	"github.com/PureMLHQ/PureML/packages/purebackend/model/middlewares"
	orgmiddlewares "github.com/PureMLHQ/PureML/packages/purebackend/user_org/middlewares"
	"github.com/labstack/echo/v4"
)

// BindModelLogsApi registers the admin api endpoints and the corresponding handlers.
func BindModelLogsApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	modelGroup := rg.Group("/org/:orgId/model", authmiddlewares.RequireAuthContext, orgmiddlewares.ValidateOrg(api.app))
	modelGroup.GET("/:modelName/branch/:branchName/version/:version/log", api.DefaultHandler(GetAllLogsModel), middlewares.ValidateModel(api.app), middlewares.ValidateModelBranch(api.app), middlewares.ValidateModelBranchVersion(api.app))
	modelGroup.GET("/:modelName/branch/:branchName/version/:version/log/:key", api.DefaultHandler(GetKeyLogsModel), middlewares.ValidateModel(api.app), middlewares.ValidateModelBranch(api.app), middlewares.ValidateModelBranchVersion(api.app))
	modelGroup.POST("/:modelName/branch/:branchName/version/:version/log", api.DefaultHandler(LogModel), middlewares.ValidateModel(api.app), middlewares.ValidateModelBranch(api.app), middlewares.ValidateModelBranchVersion(api.app))
}

// LogModel godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Log data for model
//	@Description	Log data for model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName}/version/{version}/log [post]
//	@Param			orgId		path	string				true	"Organization Id"
//	@Param			modelName	path	string				true	"Model Name"
//	@Param			branchName	path	string				true	"Branch Name"
//	@Param			version		path	string				true	"Version"
//	@Param			data		body	models.LogRequest	true	"Data to log"
func (api *Api) LogModel(request *models.Request) *models.Response {
	request.ParseJsonBody()
	key := request.GetParsedBodyAttribute("key")
	var keyData string
	if key != nil {
		keyData = key.(string)
	} else {
		keyData = ""
	}
	data := request.GetParsedBodyAttribute("data")
	var dataData string
	if data != nil {
		dataData = data.(string)
	} else {
		dataData = ""
	}
	versionUUID := request.GetModelBranchVersionUUID()
	result, err := api.app.Dao().CreateLogForModelVersion(keyData, dataData, versionUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Log created")
	return response
}

// GetAllLogsModel godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get Log data for model
//	@Description	Get Log data for model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName}/version/{version}/log [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
//	@Param			branchName	path	string	true	"Branch Name"
//	@Param			version		path	string	true	"Version"
func (api *Api) GetAllLogsModel(request *models.Request) *models.Response {
	versionUUID := request.GetModelBranchVersionUUID()
	result, err := api.app.Dao().GetLogForModelVersion(versionUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Logs for model version")
	return response
}

// GetKeyLogsModel godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get Log data for model with specific key
//	@Description	Get Log data for model with specific key
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName}/version/{version}/log/{key} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
//	@Param			branchName	path	string	true	"Branch Name"
//	@Param			version		path	string	true	"Version"
//	@Param			key			path	string	true	"Key"
func (api *Api) GetKeyLogsModel(request *models.Request) *models.Response {
	versionUUID := request.GetModelBranchVersionUUID()
	key := request.PathParams["key"]
	result, err := api.app.Dao().GetKeyLogForModelVersion(versionUUID, key)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Specific Key Logs for model version")
	return response
}

var LogModel ServiceFunc = (*Api).LogModel
var GetAllLogsModel ServiceFunc = (*Api).GetAllLogsModel
var GetKeyLogsModel ServiceFunc = (*Api).GetKeyLogsModel
