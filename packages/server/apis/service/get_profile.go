package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetProfile godoc
// @Security ApiKeyAuth
// @Summary Get logged in user profile.
// @Description Get logged in user profile.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/profile [get]
func GetProfile(request *models.Request) *models.Response {
	userUUID := request.GetUserUUID()
	user, err := datastore.GetUserByUUID(userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, user, "User profile")
}
