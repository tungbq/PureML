package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func CreateOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgName := request.GetParsedBodyAttribute("name").(string)
	email := request.User.Email
	org, err := datastore.CreateOrgFromEmail(email, orgName)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response := &models.Response{}
	response.StatusCode = http.StatusOK
	response.Body.Status = response.StatusCode
	response.Body.Message = "Organization created"
	response.Body.Data = []models.Organization{*org}
	return response
}
