package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// UserLogin godoc
// @Security ApiKeyAuth
// @Summary User login.
// @Description User login with email and password.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/login [post]
// @Param org body models.UserLoginRequest true "User details"
func UserLogin(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	handle := request.GetParsedBodyAttribute("handle")
	if email == nil && handle == nil {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Email or handle is required")
	}
	password := request.GetParsedBodyAttribute("password").(string)
	var user *models.UserResponse
	var err error
	if email != nil {
		email := email.(string)
		user, err = datastore.GetUserByEmail(email)
	} else {
		handle := handle.(string)
		user, err = datastore.GetUserByHandle(handle)
	}
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.NewDataResponse(http.StatusUnauthorized, nil, "Invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":   user.UUID,
		"email":  user.Email,
		"handle": user.Handle,
	})
	signedString, err := token.SignedString(config.TokenSigningSecret())
	if err != nil {
		panic(err)
	}
	data := []map[string]string{
		{
			"email":       user.Email,
			"accessToken": signedString,
		},
	}
	return models.NewDataResponse(http.StatusAccepted, data, "User logged in")
}
