package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetOrgsForUser(request *models.Request) *models.Response {
	email := request.User.Email
	response := &models.Response{}
	UserOrganization, err := datastore.GetUserOrganizationsByEmail(email)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response.StatusCode = http.StatusAccepted
	response.Body.Status = response.StatusCode
	response.Body.Message = "User Organizations"
	response.Body.Data = UserOrganization
	return response
}
