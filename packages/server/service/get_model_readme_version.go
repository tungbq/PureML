package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetModelReadmeVersion godoc
// @Security ApiKeyAuth
// @Summary Get model readme
// @Description Get model readme
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/readme/version/{version} [get]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param version path string true "Version"
func GetModelReadmeVersion(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	versionName := request.GetPathParam("version")
	readme, err := datastore.GetModelReadmeVersion(modelUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if readme == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Model readme version not found")
	}
	response := models.NewDataResponse(http.StatusOK, readme, "Model Readme version")
	return response
}
