package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)


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
func GetModelAllBranches(request *models.Request) *models.Response {
	var response *models.Response
	modelUUID := request.GetModelUUID()
	allOrgs, err := datastore.GetModelAllBranches(modelUUID)
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
func GetModelBranch(request *models.Request) *models.Response {
	modelBranchUUID := request.GetModelBranchUUID()
	branch, err := datastore.GetModelBranchByUUID(modelBranchUUID)
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
func CreateModelBranch(request *models.Request) *models.Response {
	request.ParseJsonBody()
	modelUUID := request.GetModelUUID()
	modelBranchName := request.GetParsedBodyAttribute("branch_name").(string)
	if modelBranchName == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Branch name cannot be empty")
	}
	modelBranches, err := datastore.GetModelAllBranches(modelUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	for _, branch := range modelBranches {
		if branch.Name == modelBranchName {
			return models.NewErrorResponse(http.StatusBadRequest, "Branch already exists")
		}
	}
	modelBranch, err := datastore.CreateModelBranch(modelUUID, modelBranchName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return models.NewDataResponse(http.StatusOK, modelBranch, "Model branch created")
}

// TODO: UpdateModelBranch godoc
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
func UpdateModelBranch(request *models.Request) *models.Response {
	return nil
}
// TODO: DeleteModelBranch godoc
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
func DeleteModelBranch(request *models.Request) *models.Response {
	return nil
}

