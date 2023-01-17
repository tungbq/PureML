package service

import "github.com/PureML-Inc/PureML/server/models"

// UpdateModelBranch godoc
// @Security ApiKeyAuth
// @Summary Update a branch of a model
// @Description Update a branch of a model
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/branch/{branchName}/update [post]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param branchName path string true "Branch Name"
func UpdateModelBranch(request *models.Request) *models.Response {
	return nil
}
