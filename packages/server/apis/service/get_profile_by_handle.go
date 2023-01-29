package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetProfileByHandle godoc
// @Summary Get user profile by handle.
// @Description Get user profile by handle. Accessible without login.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/profile/{userHandle} [get]
// @Param userHandle path string true "User handle"
func GetProfileByHandle(request *models.Request) *models.Response {
	userHandle := request.GetPathParam("userHandle")
	user, err := datastore.GetUserByHandle(userHandle)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, user, "Public User profile")
}
