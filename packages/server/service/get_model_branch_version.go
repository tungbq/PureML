package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetModelBranchVersion godoc
// @Security ApiKeyAuth
// @Summary Get specific branch version of a model
// @Description Get specific branch version of a model
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/branch/{branchName}/version/{version} [get]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param branchName path string true "Branch Name"
// @Param version path string true "Version"
func GetModelBranchVersion(request *models.Request) *models.Response {
	branchUUID := request.ModelBranch.UUID
	versionName := request.PathParams["version"]
	version, err := datastore.GetModelBranchVersion(branchUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if version == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Version not found")
	}
	return models.NewDataResponse(http.StatusOK, version, "Model branch details")
}
