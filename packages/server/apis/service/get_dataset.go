package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

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
func GetDataset(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	datasetName := request.GetDatasetName()
	dataset, err := datastore.GetDatasetByName(orgId, datasetName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if dataset == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Dataset not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.DatasetResponse{*dataset}, "Dataset successfully retrieved")
}
