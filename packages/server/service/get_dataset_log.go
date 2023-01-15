package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetLogDataset godoc
// @Security ApiKeyAuth
// @Summary Get Log data for dataset
// @Description Get Log data for dataset
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/branch/{branchName}/version/{version}/log [get]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param branchName path string true "Branch Name"
// @Param version path string true "Version"
func GetLogDataset(request *models.Request) *models.Response {
	branchUUID := request.GetDatasetBranchUUID()
	versionName := request.PathParams["version"]
	version, err := datastore.GetDatasetBranchVersion(branchUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if version == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Version not found")
	}
	result, err := datastore.GetLogForDatasetVersion(version.UUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Log created")
	return response
}
