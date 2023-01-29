package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// UpdateDatasetReadme godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update readme of a dataset for a category
//	@Description	Update readme of a dataset for a category
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/readme [post]
//	@Param			orgId		path	string					true	"Organization Id"
//	@Param			datasetName	path	string					true	"Dataset Name"
//	@Param			data		body	models.ReadmeRequest	true	"Data"
func UpdateDatasetReadme(request *models.Request) *models.Response {
	request.ParseJsonBody()
	datasetUUID := request.GetDatasetUUID()
	datasetFileType := request.GetParsedBodyAttribute("file_type")
	var datasetFileTypeData string
	if datasetFileType == nil {
		datasetFileTypeData = ""
	} else {
		datasetFileTypeData = datasetFileType.(string)
	}
	datasetContent := request.GetParsedBodyAttribute("content")
	var datasetContentData string
	if datasetContent == nil {
		datasetContentData = ""
	} else {
		datasetContentData = datasetContent.(string)
	}
	readme, err := datastore.UpdateDatasetReadme(datasetUUID, datasetFileTypeData, datasetContentData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, readme, "Dataset readme updated")
}
