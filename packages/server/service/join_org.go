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
		response.Message = "Organization not found"
		return response
	}
	mailId := request.GetUserMail()
	_, err = datastore.CreateOrgAcessFromMailIdAndJoinCode(mailId, joinCode)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response.StatusCode = http.StatusOK
	response.Message = "User joined organization"
	return response

}
