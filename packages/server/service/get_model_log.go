package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetLogModel godoc
// @Security ApiKeyAuth
// @Summary Get Log data for model
// @Description Get Log data for model
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/branch/{branchName}/version/{version}/log [get]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param branchName path string true "Branch Name"
// @Param version path string true "Version"
func GetLogModel(request *models.Request) *models.Response {
	branchUUID := request.GetModelBranchUUID()
	versionName := request.PathParams["version"]
	version, err := datastore.GetModelBranchVersion(branchUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if version == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Version not found")
	}
	result, err := datastore.GetLogForModelVersion(version.UUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Log created")
	return response
}
