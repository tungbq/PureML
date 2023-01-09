package service

import "github.com/PureML-Inc/PureML/server/models"

// GetProfileByHandle godoc
// @Summary Get user profile by handle.
// @Description Get user profile by handle. Accessible without login.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/profile/:userhandle [post]
// @Param userHandle path string true "User handle"
func GetProfileByHandle(request *models.Request) *models.Response {
	return nil
}
