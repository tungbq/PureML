package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// CreateModelActivity godoc
// @Security ApiKeyAuth
// @Summary Add activity of a model for a category
// @Description Add activity of a model for a category
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/model/{modelName}/activity/{category} [post]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param category path string true "Category"
// @Param data body models.ActivityRequest true "Activity"
func CreateModelActivity(request *models.Request) *models.Response {
	request.ParseJsonBody()
	modelUUID := request.GetModelUUID()
	userUUID := request.GetUserUUID()
	category := request.GetPathParam("category")
	activity := request.GetParsedBodyAttribute("activity")
	if activity == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Activity not found in request body")
	} else if activity.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Activity cannot be empty")
	}
	activityData := activity.(string)
	createdActivity, err := datastore.CreateModelActivity(modelUUID, userUUID, category, activityData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.ActivityResponse{*createdActivity}, "Activity created")
}
