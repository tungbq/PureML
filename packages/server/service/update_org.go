package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func UpdateOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetPathParam("orgId")
	orgName := request.GetParsedBodyAttribute("name").(string)
	email := request.User.Email
	UserOrganization, err := datastore.GetUserOrganizationByOrgIdAndEmail(orgId, email)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response := &models.Response{}
	if UserOrganization.Role != "owner" {
		response.StatusCode = http.StatusForbidden
		response.Body.Status = response.StatusCode
		response.Body.Message = "You are not authorized to update this organization"
		response.Body.Data = nil
	}
	updatedOrg, err := datastore.UpdateOrg(orgId, orgName)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response.StatusCode = http.StatusOK
	response.Body.Status = response.StatusCode
	response.Body.Message = "Organization updated"
	response.Body.Data = []models.Organization{*updatedOrg}
	return response

}
