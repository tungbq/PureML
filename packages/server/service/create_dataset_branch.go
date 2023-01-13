package service

import "github.com/PureML-Inc/PureML/server/models"

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
// @Param branchName body string true "Branch Name"
func CreateDatasetBranch(request *models.Request) *models.Response {
	return nil
}
