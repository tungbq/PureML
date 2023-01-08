package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/teris-io/shortid"
	"golang.org/x/crypto/bcrypt"
)

func UserSignUp(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	name := request.GetParsedBodyAttribute("name")
	if name == nil {
		name = ""
	}
	password := request.GetParsedBodyAttribute("password").(string)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	shortId, _ := shortid.Generate()
	user, err := datastore.CreateUser(name.(string), email, string(hashedPassword), shortId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := &models.Response{}
	response.StatusCode = http.StatusOK
	response.Body.Message = "User created"
	response.Body.Data = []models.UserResponse{*user}
	return response
}
