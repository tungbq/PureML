package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/middlewares"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
)

// BindDatasetLogsApi registers the admin api endpoints and the corresponding handlers.
func BindDatasetLogsApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	datasetGroup := rg.Group("/org/:orgId/dataset", middlewares.AuthenticateJWT, middlewares.ValidateOrg)
	datasetGroup.GET("/:datasetName/branch/:branchName/version/:version/log", api.DefaultHandler(GetAllLogsDataset), middlewares.ValidateDataset, middlewares.ValidateDatasetBranch, middlewares.ValidateDatasetBranchVersion)
	datasetGroup.GET("/:datasetName/branch/:branchName/version/:version/log/:key", api.DefaultHandler(GetKeyLogsDataset), middlewares.ValidateDataset, middlewares.ValidateDatasetBranch, middlewares.ValidateDatasetBranchVersion)
	datasetGroup.POST("/:datasetName/branch/:branchName/version/:version/log", api.DefaultHandler(LogDataset), middlewares.ValidateDataset, middlewares.ValidateDatasetBranch, middlewares.ValidateDatasetBranchVersion)
}

// LogDataset godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Log data for dataset
//	@Description	Log data for dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName}/version/{version}/log [post]
//	@Param			orgId		path	string				true	"Organization Id"
//	@Param			datasetName	path	string				true	"Dataset Name"
//	@Param			branchName	path	string				true	"Branch Name"
//	@Param			version		path	string				true	"Version"
//	@Param			data		body	models.LogRequest	true	"Data to log"
func (api *Api) LogDataset(request *models.Request) *models.Response {
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
	versionUUID := request.GetDatasetBranchVersionUUID()
	result, err := api.app.Dao().CreateLogForDatasetVersion(keyData, dataData, versionUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Log created")
	return response
}

// GetAllLogsDataset godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get Log data for dataset
//	@Description	Get Log data for dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName}/version/{version}/log [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			branchName	path	string	true	"Branch Name"
//	@Param			version		path	string	true	"Version"
func (api *Api) GetAllLogsDataset(request *models.Request) *models.Response {
	versionUUID := request.GetDatasetBranchVersionUUID()
	result, err := api.app.Dao().GetLogForDatasetVersion(versionUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Logs for dataset version")
	return response
}

// GetKeyLogsDataset godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get Log data for dataset with specific key
//	@Description	Get Log data for dataset with specific key
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName}/version/{version}/log/{key} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			branchName	path	string	true	"Branch Name"
//	@Param			version		path	string	true	"Version"
//	@Param			key			path	string	true	"Key"
func (api *Api) GetKeyLogsDataset(request *models.Request) *models.Response {
	versionUUID := request.GetDatasetBranchVersionUUID()
	key := request.PathParams["key"]
	result, err := api.app.Dao().GetKeyLogForDatasetVersion(versionUUID, key)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Specific Key Logs for dataset version")
	return response
}

var LogDataset ServiceFunc = (*Api).LogDataset
var GetAllLogsDataset ServiceFunc = (*Api).GetAllLogsDataset
var GetKeyLogsDataset ServiceFunc = (*Api).GetKeyLogsDataset
