package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// CreateDatasetBranch godoc
// @Security ApiKeyAuth
// @Summary Create a new branch of a dataset
// @Description Create a new branch of a dataset
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/branch/create [post]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param branchName body models.CreateDatasetBranchRequest true "Data"
func CreateDatasetBranch(request *models.Request) *models.Response {
	request.ParseJsonBody()
	datasetUUID := request.GetDatasetUUID()
	datasetBranchName := request.GetParsedBodyAttribute("branch_name").(string)
	if datasetBranchName == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Branch name cannot be empty")
	}
	datasetBranches, err := datastore.GetDatasetAllBranches(datasetUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	for _, branch := range datasetBranches {
		if branch.Name == datasetBranchName {
			return models.NewErrorResponse(http.StatusBadRequest, "Branch already exists")
		}
	}
	modelBranch, err := datastore.CreateDatasetBranch(datasetUUID, datasetBranchName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return models.NewDataResponse(http.StatusOK, modelBranch, "Dataset branch created")
}
