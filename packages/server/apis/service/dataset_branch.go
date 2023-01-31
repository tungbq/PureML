package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetDatasetAllBranches godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all branches of a dataset
//	@Description	Get all branches of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
func GetDatasetAllBranches(request *models.Request) *models.Response {
	var response *models.Response
	datasetUUID := request.GetDatasetUUID()
	allOrgs, err := datastore.GetDatasetAllBranches(datasetUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	} else {
		response = models.NewDataResponse(http.StatusOK, allOrgs, "All dataset branches")
	}
	return response
}

// GetDatasetBranch godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get specific branch of a dataset
//	@Description	Get specific branch of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			branchName	path	string	true	"Branch Name"
func GetDatasetBranch(request *models.Request) *models.Response {
	datasetBranchUUID := request.GetDatasetBranchUUID()
	branch, err := datastore.GetDatasetBranchByUUID(datasetBranchUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if branch == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Branch not found")
	}
	return models.NewDataResponse(http.StatusOK, branch, "Dataset branch details")
}

// CreateDatasetBranch godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create a new branch of a dataset
//	@Description	Create a new branch of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/create [post]
//	@Param			orgId		path	string								true	"Organization Id"
//	@Param			datasetName	path	string								true	"Dataset Name"
//	@Param			branchName	body	models.CreateDatasetBranchRequest	true	"Data"
func CreateDatasetBranch(request *models.Request) *models.Response {
	request.ParseJsonBody()
	datasetUUID := request.GetDatasetUUID()
	datasetBranchName := request.GetParsedBodyAttribute("branch_name").(string)
	if datasetBranchName == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Branch name cannot be empty")
	}
	datasetBranches, err := datastore.GetDatasetAllBranches(datasetUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	for _, branch := range datasetBranches {
		if branch.Name == datasetBranchName {
			return models.NewErrorResponse(http.StatusBadRequest, "Branch already exists")
		}
	}
	modelBranch, err := datastore.CreateDatasetBranch(datasetUUID, datasetBranchName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return models.NewDataResponse(http.StatusOK, modelBranch, "Dataset branch created")
}

// TODO: UpdateDatasetBranch godoc
//	@Security		ApiKeyAuth
//	@Summary		Update a branch of a dataset
//	@Description	Update a branch of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName}/update [post]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			branchName	path	string	true	"Branch Name"
func UpdateDatasetBranch(request *models.Request) *models.Response {
	return nil
}

// TODO: DeleteDatasetBranch godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete a branch of a dataset
//	@Description	Delete a branch of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName}/delete [delete]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			branchName	path	string	true	"Branch Name"
func DeleteDatasetBranch(request *models.Request) *models.Response {
	return nil
}
