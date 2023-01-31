package service

import (
	_ "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

var defaultModelBranchNames = []string{"main", "development"}

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
func GetAllModels(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	allModels, err := datastore.GetAllModels(orgId)
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
func GetModel(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	modelName := request.GetModelName()
	model, err := datastore.GetModelByName(orgId, modelName)
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
func CreateModel(request *models.Request) *models.Response {
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
		for _, branchName := range modelBranchNames {
			modelBranchNamesData = append(modelBranchNamesData, branchName.(string))
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
	model, err := datastore.GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Model already exists")
	}
	model, err = datastore.CreateModel(orgId, modelName, modelWikiData, modelIsPublicData, modelReadmeData, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	_, err = datastore.CreateModelBranches(model.UUID, modelBranchNamesData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, model, "Model and branches successfully created")
}
