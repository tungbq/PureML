package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetDatasetBranchVersion godoc
// @Security ApiKeyAuth
// @Summary Get specific branch version of a dataset
// @Description Get specific branch version of a dataset
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/branch/{branchName}/version/{version} [get]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param branchName path string true "Branch Name"
// @Param version path string true "Version"
func GetDatasetBranchVersion(request *models.Request) *models.Response {
	branchUUID := request.GetDatasetBranchUUID()
	versionName := request.PathParams["version"]
	version, err := datastore.GetDatasetBranchVersion(branchUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if version == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Version not found")
	}
	return models.NewDataResponse(http.StatusOK, version, "Dataset branch details")
}
