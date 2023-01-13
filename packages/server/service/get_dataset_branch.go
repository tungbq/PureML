package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetDatasetBranch godoc
// @Security ApiKeyAuth
// @Summary Get specific branch of a dataset
// @Description Get specific branch of a dataset
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/branch/{branchName} [get]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param branchName path string true "Branch Name"
func GetDatasetBranch(request *models.Request) *models.Response {
	datasetBranchUUID := request.DatasetBranch.UUID
	branch, err := datastore.GetModelBranchByUUID(datasetBranchUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return models.NewDataResponse(http.StatusOK, branch, "Dataset branch details")
}
