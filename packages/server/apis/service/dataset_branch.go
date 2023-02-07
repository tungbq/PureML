package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/middlewares"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
)

// BindDatasetBranchApi registers the admin api endpoints and the corresponding handlers.
func BindDatasetBranchApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	datasetGroup := rg.Group("/org/:orgId/dataset", middlewares.AuthenticateJWT(api.app), middlewares.ValidateOrg(api.app))
	datasetGroup.GET("/:datasetName/branch", api.DefaultHandler(GetDatasetAllBranches), middlewares.ValidateDataset(api.app))
	datasetGroup.POST("/:datasetName/branch/create", api.DefaultHandler(CreateDatasetBranch), middlewares.ValidateDataset(api.app))
	datasetGroup.GET("/:datasetName/branch/:branchName", api.DefaultHandler(GetDatasetBranch), middlewares.ValidateDataset(api.app), middlewares.ValidateDatasetBranch(api.app))
	datasetGroup.POST("/:datasetName/branch/:branchName/update", api.DefaultHandler(UpdateDatasetBranch), middlewares.ValidateDataset(api.app), middlewares.ValidateDatasetBranch(api.app))
	datasetGroup.DELETE("/:datasetName/branch/:branchName/delete", api.DefaultHandler(DeleteDatasetBranch), middlewares.ValidateDataset(api.app), middlewares.ValidateDatasetBranch(api.app))
}

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
func (api *Api) GetDatasetAllBranches(request *models.Request) *models.Response {
	var response *models.Response
	datasetUUID := request.GetDatasetUUID()
	allOrgs, err := api.app.Dao().GetDatasetAllBranches(datasetUUID)
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
func (api *Api) GetDatasetBranch(request *models.Request) *models.Response {
	datasetBranchUUID := request.GetDatasetBranchUUID()
	branch, err := api.app.Dao().GetDatasetBranchByUUID(datasetBranchUUID)
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
func (api *Api) CreateDatasetBranch(request *models.Request) *models.Response {
	request.ParseJsonBody()
	datasetUUID := request.GetDatasetUUID()
	datasetBranchName := request.GetParsedBodyAttribute("branch_name").(string)
	if datasetBranchName == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Branch name cannot be empty")
	}
	datasetBranches, err := api.app.Dao().GetDatasetAllBranches(datasetUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	for _, branch := range datasetBranches {
		if branch.Name == datasetBranchName {
			return models.NewErrorResponse(http.StatusBadRequest, "Branch already exists")
		}
	}
	modelBranch, err := api.app.Dao().CreateDatasetBranch(datasetUUID, datasetBranchName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return models.NewDataResponse(http.StatusOK, modelBranch, "Dataset branch created")
}

// TODO: UpdateDatasetBranch godoc
//
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
func (api *Api) UpdateDatasetBranch(request *models.Request) *models.Response {
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
func (api *Api) DeleteDatasetBranch(request *models.Request) *models.Response {
	return nil
}

var GetDatasetAllBranches ServiceFunc = (*Api).GetDatasetAllBranches
var GetDatasetBranch ServiceFunc = (*Api).GetDatasetBranch
var CreateDatasetBranch ServiceFunc = (*Api).CreateDatasetBranch
var UpdateDatasetBranch ServiceFunc = (*Api).UpdateDatasetBranch
var DeleteDatasetBranch ServiceFunc = (*Api).DeleteDatasetBranch
