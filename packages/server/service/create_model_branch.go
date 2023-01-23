package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

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
	modelUUID := request.GetModelUUID()
	modelBranchName := request.GetPathParam("branchName")
	modelBranches, err := datastore.GetModelAllBranches(modelUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	for _, branch := range modelBranches {
		if branch.Name == modelBranchName {
			return models.NewErrorResponse(http.StatusBadRequest, "Branch already exists")
		}
	}
	modelBranch, err := datastore.CreateModelBranch(modelUUID, modelBranchName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return models.NewDataResponse(http.StatusOK, modelBranch, "Model branch created")
}
