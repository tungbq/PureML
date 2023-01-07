package service

import (
	"net/http"

	"github.com/PriyavKaneria/PureML/service/config"
	"github.com/PriyavKaneria/PureML/service/datastore"
	"github.com/PriyavKaneria/PureML/service/models"
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
