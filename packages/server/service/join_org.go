package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// JoinOrg godoc
// @Summary Join organization by join code.
// @Description Join organization by join code.
// @Tags Organization
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/join [post]
// @Param join_code body string true "Organization join code"
func JoinOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	joinCode := request.GetParsedBodyAttribute("join_code").(string)
	org, err := datastore.GetOrgByJoinCode(joinCode)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	var response *models.Response
	if org == nil {
		response = models.NewErrorResponse(http.StatusNotFound, "Organization not found")
		return response
	}
	email := request.GetUserMail()
	_, err = datastore.CreateUserOrganizationFromEmailAndJoinCode(email, joinCode)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response = models.NewDataResponse(http.StatusOK, nil, "User joined organization")
	return response
}
