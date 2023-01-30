package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// AddUsersToOrg godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Add a user to an organization.
//	@Description	Add a user to an organization. Only accessible by owners of the organization.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/add [post]
//	@Param			email	path	string	true	"User email"
func AddUsersToOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	orgId := uuid.Must(uuid.FromString(request.PathParams["orgId"]))
	user, err := datastore.GetUserByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	var response *models.Response
	if user == nil {
		response = models.NewErrorResponse(http.StatusNotFound, "User not found")
		return response
	}
	_, err = datastore.CreateUserOrganizationFromEmailAndOrgId(email, orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response = models.NewDataResponse(http.StatusOK, nil, "User added to organization")
	return response
}
