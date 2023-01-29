package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetModelReadmeAllVersions godoc
// @Security ApiKeyAuth
// @Summary Get model readme
// @Description Get model readme
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/model/{modelName}/readme/version [get]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
func GetModelReadmeAllVersions(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	readme, err := datastore.GetModelReadmeAllVersions(modelUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	response := models.NewDataResponse(http.StatusOK, readme, "Model Readme version")
	return response
}
