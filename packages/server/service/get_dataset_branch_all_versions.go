package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetDatasetBranchAllVersions godoc
// @Security ApiKeyAuth
// @Summary Get all branch versions of a dataset
// @Description Get all branch versions of a dataset
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/branch/{branchName}/version [get]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param branchName path string true "Branch Name"
func GetDatasetBranchAllVersions(request *models.Request) *models.Response {
	var response *models.Response
	branchUUID := request.GetDatasetBranchUUID()
	allVersions, err := datastore.GetDatasetBranchAllVersions(branchUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	} else {
		response = models.NewDataResponse(http.StatusOK, allVersions, "All organizations")
	}
	return response
}
