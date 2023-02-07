package service

import (
	_ "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/models"
)

var defaultDatasetBranchNames = []string{"main", "development"}

// GetAllPublicDatasets godoc
//
//	@Summary		Get all public datasets
//	@Description	Get all public datasets
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/public/dataset [get]
func (api *Api) GetAllPublicDatasets(request *models.Request) *models.Response {
	allDatasets, err := api.app.Dao().GetAllPublicDatasets()
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, allDatasets, "Datasets successfully retrieved")
}

// GetAllDatasets godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all datasets of an organization
//	@Description	Get all datasets of an organization
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/all [get]
//	@Param			orgId	path	string	true	"Organization Id"
func (api *Api) GetAllDatasets(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	allDatasets, err := api.app.Dao().GetAllDatasets(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, allDatasets, "Datasets successfully retrieved")
}

// GetDataset godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get specific dataset of an organization
//	@Description	Get specific dataset of an organization
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
func (api *Api) GetDataset(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	datasetName := request.GetDatasetName()
	dataset, err := api.app.Dao().GetDatasetByName(orgId, datasetName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if dataset == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Dataset not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.DatasetResponse{*dataset}, "Dataset successfully retrieved")
}

// CreateDataset godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create dataset
//	@Description	Register dataset file. Create dataset and default branches if not exists
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/create [post]
//	@Param			orgId		path	string						true	"Organization UUID"
//	@Param			datasetName	path	string						true	"Dataset name"
//	@Param			data		body	models.CreateDatasetRequest	true	"Dataset details"
func (api *Api) CreateDataset(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetOrgId()
	userUUID := request.GetUserUUID()
	datasetName := request.GetPathParam("datasetName")
	datasetWiki := request.GetParsedBodyAttribute("wiki")
	var datasetWikiData string
	if datasetWiki == nil {
		datasetWikiData = ""
	} else {
		datasetWikiData = datasetWiki.(string)
	}
	datasetIsPublic := request.GetParsedBodyAttribute("is_public")
	var datasetIsPublicData bool
	if datasetIsPublic == nil {
		datasetIsPublicData = false
	} else {
		datasetIsPublicData = datasetIsPublic.(bool)
	}
	datasetBranchNames := request.GetParsedBodyAttribute("branch_names")
	var datasetBranchNamesData []string
	if datasetBranchNames == nil {
		datasetBranchNamesData = defaultDatasetBranchNames
	} else {
		datasetBranchNames := datasetBranchNames.([]interface{})
		for _, branchName := range datasetBranchNames {
			datasetBranchNamesData = append(datasetBranchNamesData, branchName.(string))
		}
	}
	datasetReadme := request.GetParsedBodyAttribute("readme")
	var datasetReadmeData *models.ReadmeRequest
	if datasetReadme == nil {
		datasetReadmeData = &models.ReadmeRequest{
			FileType: "markdown",
			Content:  "",
		}
	} else {
		datasetReadmeData = &models.ReadmeRequest{
			FileType: datasetReadme.(map[string]interface{})["file_type"].(string),
			Content:  datasetReadme.(map[string]interface{})["content"].(string),
		}
	}
	dataset, err := api.app.Dao().GetDatasetByName(orgId, datasetName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if dataset != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Dataset already exists")
	}
	dataset, err = api.app.Dao().CreateDataset(orgId, datasetName, datasetWikiData, datasetIsPublicData, datasetReadmeData, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	_, err = api.app.Dao().CreateDatasetBranches(dataset.UUID, datasetBranchNamesData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, dataset, "Dataset and branches successfully created")
}
