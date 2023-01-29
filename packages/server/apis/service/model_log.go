package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// LogModel godoc
// @Security ApiKeyAuth
// @Summary Log data for model
// @Description Log data for model
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/model/{modelName}/branch/{branchName}/version/{version}/log [post]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param branchName path string true "Branch Name"
// @Param version path string true "Version"
// @Param data body models.LogRequest true "Data to log"
func LogModel(request *models.Request) *models.Response {
	request.ParseJsonBody()
	data := request.GetParsedBodyAttribute("data").(string)
	branchUUID := request.GetModelBranchUUID()
	versionName := request.PathParams["version"]
	version, err := datastore.GetModelBranchVersion(branchUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if version == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Version not found")
	}
	result, err := datastore.CreateLogForModelVersion(data, version.UUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Log created")
	return response
}
