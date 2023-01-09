package service

import "github.com/PureML-Inc/PureML/server/models"

// UserResetPassword godoc
// @Summary User reset password.
// @Description User can reset password by providing old password and new password.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/reset-password [post]
func UserResetPassword(request *models.Request) *models.Response {
	return nil
}
