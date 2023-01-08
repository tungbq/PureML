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
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, []models.Organization{*org}, "Organization created")
	return response
}
