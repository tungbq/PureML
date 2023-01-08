package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetAllAdminOrgs(request *models.Request) *models.Response {
	response := &models.Response{}
	if config.HasAdminAccess(request.User.MailId) {
		allOrgs, err := datastore.GetAllAdminOrgs()
		if err != nil {
			return models.NewErrorResponse(err)
		}
		response.StatusCode = http.StatusOK
		response.Body = allOrgs
	} else {
		response.StatusCode = http.StatusForbidden
		response.Body = "Forbidden"
	}
	return response
}
