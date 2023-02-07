package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/middlewares"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
)

// BindModelBranchApi registers the admin api endpoints and the corresponding handlers.
func BindModelBranchApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	modelGroup := rg.Group("/org/:orgId/model", middlewares.AuthenticateJWT(api.app), middlewares.ValidateOrg(api.app))
	modelGroup.GET("/:modelName/branch", api.DefaultHandler(GetModelAllBranches), middlewares.ValidateModel(api.app))
	modelGroup.POST("/:modelName/branch/create", api.DefaultHandler(CreateModelBranch), middlewares.ValidateModel(api.app))
	modelGroup.GET("/:modelName/branch/:branchName", api.DefaultHandler(GetModelBranch), middlewares.ValidateModel(api.app), middlewares.ValidateModelBranch(api.app))
}

// GetModelAllBranches godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all branches of a model
//	@Description	Get all branches of a model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
func (api *Api) GetModelAllBranches(request *models.Request) *models.Response {
	var response *models.Response
	modelUUID := request.GetModelUUID()
	allOrgs, err := api.app.Dao().GetModelAllBranches(modelUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	} else {
		response = models.NewDataResponse(http.StatusOK, allOrgs, "All model branches")
	}
	return response
}

// GetModelBranch godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get specific branch of a model
//	@Description	Get specific branch of a model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
//	@Param			branchName	path	string	true	"Branch Name"
func (api *Api) GetModelBranch(request *models.Request) *models.Response {
	modelBranchUUID := request.GetModelBranchUUID()
	branch, err := api.app.Dao().GetModelBranchByUUID(modelBranchUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if branch == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Branch not found")
	}
	return models.NewDataResponse(http.StatusOK, branch, "Model branch details")
}

// CreateModelBranch godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create a new branch of a model
//	@Description	Create a new branch of a model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/create [post]
//	@Param			orgId		path	string							true	"Organization Id"
//	@Param			modelName	path	string							true	"Model Name"
//	@Param			branchName	body	models.CreateModelBranchRequest	true	"Data"
func (api *Api) CreateModelBranch(request *models.Request) *models.Response {
	request.ParseJsonBody()
	modelUUID := request.GetModelUUID()
	modelBranchName := request.GetParsedBodyAttribute("branch_name").(string)
	if modelBranchName == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Branch name cannot be empty")
	}
	modelBranches, err := api.app.Dao().GetModelAllBranches(modelUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	for _, branch := range modelBranches {
		if branch.Name == modelBranchName {
			return models.NewErrorResponse(http.StatusBadRequest, "Branch already exists")
		}
	}
	modelBranch, err := api.app.Dao().CreateModelBranch(modelUUID, modelBranchName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return models.NewDataResponse(http.StatusOK, modelBranch, "Model branch created")
}

// TODO: UpdateModelBranch godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update a branch of a model
//	@Description	Update a branch of a model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName}/update [post]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
//	@Param			branchName	path	string	true	"Branch Name"
func (api *Api) UpdateModelBranch(request *models.Request) *models.Response {
	return nil
}

// TODO: DeleteModelBranch godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete a branch of a model
//	@Description	Delete a branch of a model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName}/delete [delete]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
//	@Param			branchName	path	string	true	"Branch Name"
func (api *Api) DeleteModelBranch(request *models.Request) *models.Response {
	return nil
}

var GetModelAllBranches ServiceFunc = (*Api).GetModelAllBranches
var GetModelBranch ServiceFunc = (*Api).GetModelBranch
var CreateModelBranch ServiceFunc = (*Api).CreateModelBranch
var UpdateModelBranch ServiceFunc = (*Api).UpdateModelBranch
var DeleteModelBranch ServiceFunc = (*Api).DeleteModelBranch
