package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func CreateOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgName := request.GetParsedBodyAttribute("name").(string)
	orgDesc := request.GetParsedBodyAttribute("description").(string)
	orgHandle := request.GetParsedBodyAttribute("handle").(string)
	email := request.User.Email
	org, err := datastore.CreateOrgFromEmail(email, orgName, orgDesc, orgHandle)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*org}, "Organization created")
	return response
}
