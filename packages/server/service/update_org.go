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
	orgDesc := request.GetParsedBodyAttribute("description").(string)
	orgAvatar := request.GetParsedBodyAttribute("avatar").(string)
	email := request.User.Email
	UserOrganization, err := datastore.GetUserOrganizationByOrgIdAndEmail(orgId, email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	var response *models.Response
	if UserOrganization.Role != "owner" {
		response = models.NewErrorResponse(http.StatusForbidden, "You are not authorized to update this organization")
		return response
	}
	updatedOrg, err := datastore.UpdateOrg(orgId, orgName, orgDesc, orgAvatar)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response = models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*updatedOrg}, "Organization updated")
	return response

}
