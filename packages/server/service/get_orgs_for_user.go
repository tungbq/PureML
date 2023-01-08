package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetOrgsForUser(request *models.Request) *models.Response {
	mailId := request.User.MailId
	response := &models.Response{}
	orgAccess, err := datastore.GetOrgAccessesByMailId(mailId)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response.StatusCode = http.StatusAccepted
	response.Body = orgAccess
	return response
}
