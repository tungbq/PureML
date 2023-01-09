package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// Log godoc
// @Security ApiKeyAuth
// @Summary Log data for dataset or model
// @Description Log data for dataset or model
// @Tags Common
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /log [post]
// @Param data body models.LogRequest true "Data to log"
func Log(request *models.Request) *models.Response {
	request.ParseJsonBody()
	data := request.GetParsedBodyAttribute("data").(string)
	modelVersionUUID := request.GetParsedBodyAttribute("model_version_uuid")
	datasetVersionUUID := request.GetParsedBodyAttribute("dataset_version_uuid")
	if (modelVersionUUID == nil || modelVersionUUID == "") && (datasetVersionUUID == nil || datasetVersionUUID == "") {
		return models.NewErrorResponse(http.StatusBadRequest, "Must provide model_version_uuid or dataset_version_uuid")
	}
	var err error
	var result *models.LogResponse
	if modelVersionUUID != nil && modelVersionUUID != "" {
		modelVersionUUID := uuid.Must(uuid.FromString(modelVersionUUID.(string)))
		result, err = datastore.CreateLogForModelVersion(data, modelVersionUUID)
	} else {
		datasetVersionUUID := uuid.Must(uuid.FromString(datasetVersionUUID.(string)))
		result, err = datastore.CreateLogForDatasetVersion(data, datasetVersionUUID)
	}
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Log created")
	return response
}
