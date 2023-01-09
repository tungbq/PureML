package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/golang-jwt/jwt/v4"
)

// UserForgotPassword godoc
// @Security ApiKeyAuth
// @Summary User forgot password.
// @Description User can reset password by providing email id to send reset password link.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/forgot-password [post]
// @Param org body models.UserResetPasswordRequest true "User email"
func UserForgotPassword(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	user, err := datastore.GetUserByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User with given email not found")
	}
	token := jwt.NewWithClaims(&jwt.SigningMethodHMAC{}, jwt.MapClaims{
		"email": user.Email,
	})
	_, err = token.SignedString(config.TokenSigningSecret())
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	//TODO : Invoke mailService here
	return models.NewDataResponse(http.StatusOK, nil, "Reset password link sent to your mail")
}
