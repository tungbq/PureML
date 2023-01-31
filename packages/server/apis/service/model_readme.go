package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

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
func GetModelReadmeAllVersions(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	readme, err := datastore.GetModelReadmeAllVersions(modelUUID)
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
func GetModelReadmeVersion(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	versionName := request.GetPathParam("version")
	readme, err := datastore.GetModelReadmeVersion(modelUUID, versionName)
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
func UpdateModelReadme(request *models.Request) *models.Response {
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
	readme, err := datastore.UpdateModelReadme(modelUUID, modelFileTypeData, modelContentData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, readme, "Model readme updated")
}
