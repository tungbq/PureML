package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"golang.org/x/crypto/bcrypt"
)

// UserSignUp godoc
// @Security ApiKeyAuth
// @Summary User sign up.
// @Description User sign up with email, name, handle and password.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/signup [post]
// @Param org body models.UserSignupRequest true "User details"
func UserSignUp(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	handle := request.GetParsedBodyAttribute("handle").(string)
	name := request.GetParsedBodyAttribute("name")
	if name == nil {
		name = ""
	}
	bio := request.GetParsedBodyAttribute("bio")
	if bio == nil {
		bio = ""
	}
	avatar := request.GetParsedBodyAttribute("avatar")
	if avatar == nil {
		avatar = ""
	}
	password := request.GetParsedBodyAttribute("password").(string)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	user, err := datastore.CreateUser(name.(string), email, handle, bio.(string), avatar.(string), string(hashedPassword))
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, []models.UserResponse{*user}, "User created")
	return response
}
