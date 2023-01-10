package service

import "github.com/PureML-Inc/PureML/server/models"

// GetModelBranchAllVersions godoc
// @Security ApiKeyAuth
// @Summary Get all branch versions of a model
// @Description Get all branch versions of a model
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/branch/{branchName}/version [get]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param branchName path string true "Branch Name"
func GetModelBranchAllVersions(request *models.Request) *models.Response {
	return nil
}
