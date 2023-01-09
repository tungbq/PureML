package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetProfile godoc
// @Summary Get logged in user profile.
// @Description Get logged in user profile.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/profile [get]
func GetProfile(request *models.Request) *models.Response {
	email := request.GetUserMail()
	user, err := datastore.GetUserByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := &models.Response{}
	response.StatusCode = http.StatusOK
	response.Body.Message = "User profile"
	response.Body.Data = []models.UserResponse{
		*user,
	}
	return response
}
