package service

import (
	_ "fmt"
	"net/http"
	"strings"

	"github.com/PuremlHQ/PureML/packages/purebackend/core"
	"github.com/PuremlHQ/PureML/packages/purebackend/middlewares"
	"github.com/PuremlHQ/PureML/packages/purebackend/models"
	"github.com/labstack/echo/v4"
)

var defaultModelBranchNames = []string{"main", "development"}

// BindModelApi registers the admin api endpoints and the corresponding handlers.
func BindModelApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	rg.GET("/public/model", api.DefaultHandler(GetAllPublicModels))
	modelGroup := rg.Group("/org/:orgId/model", middlewares.RequireAuthContext, middlewares.ValidateOrg(api.app))
	modelGroup.GET("/all", api.DefaultHandler(GetAllModels))
	modelGroup.GET("/:modelName", api.DefaultHandler(GetModel), middlewares.ValidateModel(api.app))
	modelGroup.POST("/:modelName/create", api.DefaultHandler(CreateModel))
}

// GetAllPublicModels godoc
//
//	@Summary		Get all public models
//	@Description	Get all public models
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/public/model [get]
func (api *Api) GetAllPublicModels(request *models.Request) *models.Response {
	allModels, err := api.app.Dao().GetAllPublicModels()
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, allModels, "Models successfully retrieved")
}

// GetAllModels godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all models of an organization
//	@Description	Get all models of an organization
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/all [get]
//	@Param			orgId	path	string	true	"Organization Id"
func (api *Api) GetAllModels(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	allModels, err := api.app.Dao().GetAllModels(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, allModels, "Models successfully retrieved")
}

// GetModel godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get specific model of an organization
//	@Description	Get specific model of an organization
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
func (api *Api) GetModel(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	modelName := request.GetModelName()
	model, err := api.app.Dao().GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Model not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.ModelResponse{*model}, "Model successfully retrieved")
}

// CreateModel godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create model
//	@Description	Register model file. Create model and default branches if not exists
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/create [post]
//	@Param			orgId		path	string						true	"Organization UUID"
//	@Param			modelName	path	string						true	"Model name"
//	@Param			data		body	models.CreateModelRequest	true	"Model details"
func (api *Api) CreateModel(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetOrgId()
	userUUID := request.GetUserUUID()
	modelName := request.GetPathParam("modelName")
	modelWiki := request.GetParsedBodyAttribute("wiki")
	var modelWikiData string
	if modelWiki == nil {
		modelWikiData = ""
	} else {
		modelWikiData = modelWiki.(string)
	}
	modelIsPublic := request.GetParsedBodyAttribute("is_public")
	var modelIsPublicData bool
	if modelIsPublic == nil {
		modelIsPublicData = false
	} else {
		modelIsPublicData = modelIsPublic.(bool)
	}
	modelBranchNames := request.GetParsedBodyAttribute("branch_names")
	var modelBranchNamesData []string
	if modelBranchNames == nil {
		modelBranchNamesData = defaultModelBranchNames
	} else {
		modelBranchNames := modelBranchNames.([]interface{})
		hasMain := false
		for _, branchName := range modelBranchNames {
			branchName := strings.ToLower(branchName.(string))
			if branchName == "main" {
				hasMain = true
			}
			modelBranchNamesData = append(modelBranchNamesData, branchName)
		}
		if !hasMain {
			return models.NewErrorResponse(http.StatusBadRequest, "Branch names must contain 'main'")
		}
	}
	modelReadme := request.GetParsedBodyAttribute("readme")
	var modelReadmeData *models.ReadmeRequest
	if modelReadme == nil {
		modelReadmeData = &models.ReadmeRequest{
			FileType: "markdown",
			Content:  "",
		}
	} else {
		modelReadmeData = &models.ReadmeRequest{
			FileType: modelReadme.(map[string]interface{})["file_type"].(string),
			Content:  modelReadme.(map[string]interface{})["content"].(string),
		}
	}
	model, err := api.app.Dao().GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Model already exists")
	}
	model, err = api.app.Dao().CreateModel(orgId, modelName, modelWikiData, modelIsPublicData, modelReadmeData, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	_, err = api.app.Dao().CreateModelBranches(model.UUID, modelBranchNamesData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, model, "Model and branches successfully created")
}

var GetAllPublicModels ServiceFunc = (*Api).GetAllPublicModels
var GetAllModels ServiceFunc = (*Api).GetAllModels
var GetModel ServiceFunc = (*Api).GetModel
var CreateModel ServiceFunc = (*Api).CreateModel
