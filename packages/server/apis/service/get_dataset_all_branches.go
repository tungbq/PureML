package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetDatasetAllBranches godoc
// @Security ApiKeyAuth
// @Summary Get all branches of a dataset
// @Description Get all branches of a dataset
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/dataset/{datasetName}/branch [get]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
func GetDatasetAllBranches(request *models.Request) *models.Response {
	var response *models.Response
	datasetUUID := request.GetDatasetUUID()
	allOrgs, err := datastore.GetDatasetAllBranches(datasetUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	} else {
		response = models.NewDataResponse(http.StatusOK, allOrgs, "All dataset branches")
	}
	return response
}
