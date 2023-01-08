package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func JoinOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	joinCode := request.GetParsedBodyAttribute("join_code").(string)
	org, err := datastore.GetOrgByJoinCode(joinCode)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response := &models.Response{}
	if org == nil {
		response.StatusCode = http.StatusNotFound
		response.Body.Status = http.StatusNotFound
		response.Body.Message = "Organization not found"
		response.Body.Data = nil
		return response
	}
	email := request.GetUserMail()
	_, err = datastore.CreateOrgAcessFromEmailAndJoinCode(email, joinCode)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response.StatusCode = http.StatusOK
	response.Body.Status = response.StatusCode
	response.Body.Message = "User joined organization"
	response.Body.Data = nil
	return response
}
