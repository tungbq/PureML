package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetOrgsForUser(request *models.Request) *models.Response {
	email := request.User.Email
	UserOrganization, err := datastore.GetUserOrganizationsByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusAccepted, UserOrganization, "User Organizations")
	return response
}
