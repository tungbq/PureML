package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetAllDatasets godoc
// @Security ApiKeyAuth
// @Summary Get all datasets of an organization
// @Description Get all datasets of an organization
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/dataset/all [get]
// @Param orgId path string true "Organization Id"
func GetAllDatasets(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	allDatasets, err := datastore.GetAllDatasets(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, allDatasets, "Datasets successfully retrieved")
}
