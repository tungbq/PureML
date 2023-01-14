package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetDatasetActivity godoc
// @Security ApiKeyAuth
// @Summary Get activity of a dataset for a category
// @Description Get activity of a dataset for a category
// @Tags Common
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/activity/{category} [get]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param category path string true "Category"
func GetDatasetActivity(request *models.Request) *models.Response {
	datasetUUID := request.GetDatasetUUID()
	category := request.GetPathParam("category")
	activity, err := datastore.GetDatasetActivity(datasetUUID, category)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if activity == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Activity not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*activity}, "Activity found")
}
