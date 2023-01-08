package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func AddOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	user, err := datastore.GetUser(email)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response := &models.Response{}
	if user == nil {
		response.StatusCode = http.StatusNotFound
		response.Body.Status = http.StatusNotFound
		response.Body.Message = "User not found"
		response.Body.Data = nil
		return response
	}
	orgId := request.GetOrgId()
	_, err = datastore.CreateUserOrganizationFromEmailAndOrgId(email, orgId)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response.StatusCode = http.StatusOK
	response.Body.Status = response.StatusCode
	response.Body.Data = nil
	response.Body.Message = "User added to organization"
	return response
}
