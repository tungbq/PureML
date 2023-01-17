package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// LogDataset godoc
// @Security ApiKeyAuth
// @Summary Log data for dataset
// @Description Log data for dataset
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/branch/{branchName}/version/{version}/log [post]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param branchName path string true "Branch Name"
// @Param version path string true "Version"
// @Param data body models.LogRequest true "Data to log"
func LogDataset(request *models.Request) *models.Response {
	request.ParseJsonBody()
	data := request.GetParsedBodyAttribute("data").(string)
	branchUUID := request.GetDatasetBranchUUID()
	versionName := request.PathParams["version"]
	version, err := datastore.GetDatasetBranchVersion(branchUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if version == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Version not found")
	}
	result, err := datastore.CreateLogForDatasetVersion(data, version.UUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "Log created")
	return response
}
