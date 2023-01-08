package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func LeaveOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	orgId := request.GetOrgId()
	_, err := datastore.DeleteUserOrganizationFromEmailAndOrgId(email, orgId)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response := &models.Response{}
	response.StatusCode = http.StatusOK
	response.Body.Status = response.StatusCode
	response.Body.Message = "User left organization"
	response.Body.Data = nil
	return response
}
