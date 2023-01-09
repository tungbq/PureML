package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/golang-jwt/jwt/v4"
)

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
