package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetDatasetReadmeVersion godoc
// @Security ApiKeyAuth
// @Summary Get dataset readme
// @Description Get dataset readme
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/readme/version/{version} [get]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param version path string true "Version"
func GetDatasetReadmeVersion(request *models.Request) *models.Response {
	modelUUID := request.GetDatasetUUID()
	versionName := request.GetPathParam("version")
	readme, err := datastore.GetDatasetReadmeVersion(modelUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if readme == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Dataset readme version not found")
	}
	response := models.NewDataResponse(http.StatusOK, readme, "Dataset Readme version")
	return response
}
