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
	mailId := request.User.MailId
	orgAccess, err := datastore.GetOrgAccessByOrgIdAndMailId(orgId, mailId)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response := &models.Response{}
	if orgAccess.Role != "owner" {
		response.StatusCode = http.StatusForbidden
		response.Message = "You are not authorized to update this organization"
	}
	updatedOrg, err := datastore.UpdateOrg(orgId, orgName)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response.StatusCode = http.StatusOK
	response.Message = "Organization updated"
	response.Body = []models.Organization{*updatedOrg}
	return response

}
