package service

import "github.com/PureML-Inc/PureML/server/models"

// DeleteProfile godoc
// @Security ApiKeyAuth
// @Summary Delete user profile.
// @Description Delete user profile.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/delete [delete]
func DeleteProfile(request *models.Request) *models.Response {
	return nil
}
