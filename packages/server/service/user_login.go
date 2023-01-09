package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	password := request.GetParsedBodyAttribute("password").(string)
	user, err := datastore.GetUserByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		token := jwt.NewWithClaims(&jwt.SigningMethodHMAC{}, jwt.MapClaims{
			"id":    user.Id,
			"email": user.Email,
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
	return models.NewDataResponse(http.StatusUnauthorized, nil, "Invalid credentials")
}
