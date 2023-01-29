package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetModelBranch godoc
// @Security ApiKeyAuth
// @Summary Get specific branch of a model
// @Description Get specific branch of a model
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/model/{modelName}/branch/{branchName} [get]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param branchName path string true "Branch Name"
func GetModelBranch(request *models.Request) *models.Response {
	modelBranchUUID := request.GetModelBranchUUID()
	branch, err := datastore.GetModelBranchByUUID(modelBranchUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if branch == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Branch not found")
	}
	return models.NewDataResponse(http.StatusOK, branch, "Model branch details")
}
