package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetModelActivity godoc
// @Security ApiKeyAuth
// @Summary Get activity of a model for a category
// @Description Get activity of a model for a category
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/activity/{category} [get]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param category path string true "Category"
func GetModelActivity(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	category := request.GetPathParam("category")
	activity, err := datastore.GetModelActivity(modelUUID, category)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if activity == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "Activity not found")
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*activity}, "Activity found")
}
