package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)


// GetOrgsForUser godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all user organizations.
//	@Description	Get all user organizations.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/ [get]
func GetOrgsForUser(request *models.Request) *models.Response {
	email := request.GetUserMail()
	UserOrganization, err := datastore.GetUserOrganizationsByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusAccepted, UserOrganization, "User Organizations")
	return response
}


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

// JoinOrg godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Join organization by join code.
//	@Description	Join organization by join code.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/join [post]
//	@Param			join_code	body	string	true	"Organization join code"
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

// LeaveOrg godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Leave organization.
//	@Description	Leave organization.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/leave [post]
//	@Param			orgId	path	string	true	"Organization ID"
//	@Param			email	body	string	true	"User email"
func LeaveOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	orgId := request.GetOrgId()
	err := datastore.DeleteUserOrganizationFromEmailAndOrgId(email, orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, nil, "User left organization")
	return response
}

// TODO:RemoveOrg godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Remove user from organization.
//	@Description	Remove user from organization.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/remove [post]
//	@Param			orgId	path	string	true	"Organization ID"
//	@Param			email	body	string	true	"User email"
func RemoveOrg(request *models.Request) *models.Response {
	return nil
}
