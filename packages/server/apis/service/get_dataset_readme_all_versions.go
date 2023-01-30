package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetDatasetReadmeAllVersions godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get dataset readme
//	@Description	Get dataset readme
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/readme/version [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
func GetDatasetReadmeAllVersions(request *models.Request) *models.Response {
	modelUUID := request.GetDatasetUUID()
	readme, err := datastore.GetDatasetReadmeAllVersions(modelUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	response := models.NewDataResponse(http.StatusOK, readme, "Dataset Readme version")
	return response
}
