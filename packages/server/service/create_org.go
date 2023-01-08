package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func CreateOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgName := request.GetParsedBodyAttribute("name").(string)
	mailId := request.User.MailId
	org, err := datastore.CreateOrgFromMailId(mailId, orgName)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response := &models.Response{}
	response.StatusCode = http.StatusOK
	response.Body = []models.Organization{*org}
	return response
}
