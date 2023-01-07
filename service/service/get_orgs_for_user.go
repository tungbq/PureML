package service

import (
	"net/http"

	"github.com/PriyavKaneria/PureML/service/datastore"
	"github.com/PriyavKaneria/PureML/service/models"
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
