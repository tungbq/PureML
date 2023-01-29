package service

import "github.com/PureML-Inc/PureML/server/models"

// UpdateDatasetBranch godoc
// @Security ApiKeyAuth
// @Summary Update a branch of a dataset
// @Description Update a branch of a dataset
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/dataset/{datasetName}/branch/{branchName}/update [post]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param branchName path string true "Branch Name"
func UpdateDatasetBranch(request *models.Request) *models.Response {
	return nil
}
