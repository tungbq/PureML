package service

import "github.com/PureML-Inc/PureML/server/models"

// DeleteModelBranch godoc
// @Security ApiKeyAuth
// @Summary Delete a branch of a model
// @Description Delete a branch of a model
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/:orgId/model/:modelName/branch/:branchName/delete [delete]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param branchName path string true "Branch Name"
func DeleteModelBranch(request *models.Request) *models.Response {
	return nil
}
