package service

import "github.com/PureML-Inc/PureML/server/models"

// CreateModelBranch godoc
// @Security ApiKeyAuth
// @Summary Create a new branch of a model
// @Description Create a new branch of a model
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/branch/create [post]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param branchName body string true "Branch Name"
func CreateModelBranch(request *models.Request) *models.Response {
	return nil
}
