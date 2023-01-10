package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// LogModel godoc
// @Security ApiKeyAuth
// @Summary Log data for model
// @Description Log data for model
// @Tags Common
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /model/:modelName/log [post]
// @Param data body models.LogRequest true "Data to log"
func LogModel(request *models.Request) *models.Response {
	request.ParseJsonBody()
	data := request.GetParsedBodyAttribute("data").(string)
	modelVersionUUID := request.GetModelUUID()
	result, err := datastore.CreateLogForModelVersion(data, modelVersionUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Log created")
	return response
}

// LogDataset godoc
// @Security ApiKeyAuth
// @Summary Log data for dataset
// @Description Log data for dataset
// @Tags Common
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /dataset/:datasetName/log [post]
// @Param data body models.LogRequest true "Data to log"
func LogDataset(request *models.Request) *models.Response {
	request.ParseJsonBody()
	data := request.GetParsedBodyAttribute("data").(string)
	datasetVersionUUID := request.GetDatasetUUID()
	result, err := datastore.CreateLogForDatasetVersion(data, datasetVersionUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Log created")
	return response
}
