package service

import (
	"net/http"

	authmiddlewares "github.com/PureMLHQ/PureML/packages/purebackend/auth/middlewares"
	"github.com/PureMLHQ/PureML/packages/purebackend/core"
	"github.com/PureMLHQ/PureML/packages/purebackend/core/models"
	coreservice "github.com/PureMLHQ/PureML/packages/purebackend/core/apis/service"
	orgmiddlewares "github.com/PureMLHQ/PureML/packages/purebackend/user_org/middlewares"
	"github.com/labstack/echo/v4"
)

// BindUserOrgApi registers the admin api endpoints and the corresponding handlers.
func BindUserOrgApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	orgGroup := rg.Group("/org", authmiddlewares.RequireAuthContext)
	orgGroup.GET("", api.DefaultHandler(GetOrgsForUser))
	orgGroup.POST("/:orgId/add", api.DefaultHandler(AddUsersToOrg), orgmiddlewares.ValidateOrg(api.app))
	orgGroup.POST("/:orgId/role", api.DefaultHandler(UpdateUserRole), orgmiddlewares.ValidateOrg(api.app))
	orgGroup.POST("/join", api.DefaultHandler(JoinOrg))
	orgGroup.POST("/:orgId/remove", api.DefaultHandler(RemoveOrg), orgmiddlewares.ValidateOrg(api.app))
	orgGroup.GET("/:orgId/leave", api.DefaultHandler(LeaveOrg), orgmiddlewares.ValidateOrg(api.app))
}

// GetOrgsForUser godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all user organizations.
//	@Description	Get all user organizations.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org [get]
func (api *Api) GetOrgsForUser(request *models.Request) *models.Response {
	email := request.GetUserMail()
	userOrganization, err := api.app.Dao().GetUserOrganizationsByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, userOrganization, "User Organizations")
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
//	@Param			orgId	path	string					true	"Organization ID"
//	@Param			data	body	models.UserEmailRequest	true	"User email to add"
func (api *Api) AddUsersToOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	var emailData string
	if email == nil {
		emailData = ""
	} else {
		emailData = email.(string)
	}
	if emailData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is required")
	}
	if addr, ok := coreservice.ValidateMailAddress(emailData); ok {
		emailData = addr
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is invalid")
	}
	orgId := request.GetOrgId()
	user, err := api.app.Dao().GetUserByEmail(emailData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	var response *models.Response
	if user == nil {
		response = models.NewErrorResponse(http.StatusNotFound, "User to add not found")
		return response
	}
	userUUID := request.GetUserUUID()
	userOrganization, err := api.app.Dao().GetUserOrganizationByOrgIdAndUserUUID(orgId, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if userOrganization == nil || userOrganization.Role != "owner" {
		return models.NewErrorResponse(http.StatusForbidden, "You are not authorized to add users this organization")
	}
	userOrganization, err = api.app.Dao().GetUserOrganizationByOrgIdAndUserUUID(orgId, user.UUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if userOrganization != nil {
		return models.NewErrorResponse(http.StatusConflict, "User already added to organization")
	}
	_, err = api.app.Dao().CreateUserOrganizationFromEmailAndOrgId(emailData, orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response = models.NewDataResponse(http.StatusOK, nil, "User added to organization")
	return response
}

// UpdateUserRole godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update a user's role in an organization.
//	@Description	Update a user's role in an organization. Only accessible by owners of the organization.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/role [post]
//	@Param			orgId	path	string					true	"Organization ID"
//	@Param			data	body	models.UserRoleRequest	true	"User email and role to update"
func (api *Api) UpdateUserRole(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	var emailData string
	if email == nil {
		emailData = ""
	} else {
		emailData = email.(string)
	}
	if emailData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is required")
	}
	if addr, ok := coreservice.ValidateMailAddress(emailData); ok {
		emailData = addr
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is invalid")
	}
	role := request.GetParsedBodyAttribute("role")
	var roleData string
	if role == nil {
		roleData = ""
	} else {
		roleData = role.(string)
	}
	if roleData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Role is required")
	}
	if roleData != "owner" && roleData != "member" {
		return models.NewErrorResponse(http.StatusBadRequest, "Role must be one of 'owner' or 'member'")
	}
	orgId := request.GetOrgId()
	user, err := api.app.Dao().GetUserByEmail(emailData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewErrorResponse(http.StatusNotFound, "User to update not found")
	}
	userUUID := request.GetUserUUID()
	userOrganization, err := api.app.Dao().GetUserOrganizationByOrgIdAndUserUUID(orgId, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if userOrganization == nil || userOrganization.Role != "owner" {
		return models.NewErrorResponse(http.StatusForbidden, "You are not authorized to update users in this organization")
	}
	userOrganization, err = api.app.Dao().GetUserOrganizationByOrgIdAndUserUUID(orgId, user.UUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if userOrganization == nil {
		return models.NewErrorResponse(http.StatusNotFound, "User not member of organization")
	}
	err = api.app.Dao().UpdateUserRoleByOrgIdAndUserUUID(orgId, user.UUID, roleData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "User role updated")
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
//	@Param			data	body	models.UserOrgJoin	true	"Organization join code"
func (api *Api) JoinOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	joinCode := request.GetParsedBodyAttribute("join_code")
	var joinCodeData string
	if joinCode == nil {
		joinCodeData = ""
	} else {
		joinCodeData = joinCode.(string)
	}
	if joinCodeData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Join code is required")
	}
	org, err := api.app.Dao().GetOrgByJoinCode(joinCodeData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	var response *models.Response
	if org == nil {
		response = models.NewErrorResponse(http.StatusNotFound, "Invalid join code")
		return response
	}
	email := request.GetUserMail()
	userUUID := request.GetUserUUID()
	userOrganization, err := api.app.Dao().GetUserOrganizationByOrgIdAndUserUUID(org.UUID, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if userOrganization != nil {
		return models.NewErrorResponse(http.StatusConflict, "User already member of organization")
	}
	_, err = api.app.Dao().CreateUserOrganizationFromEmailAndJoinCode(email, joinCodeData)
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
//	@Router			/org/{orgId}/leave [get]
//	@Param			orgId	path	string	true	"Organization ID"
func (api *Api) LeaveOrg(request *models.Request) *models.Response {
	email := request.GetUserMail()
	orgId := request.GetOrgId()
	userUUID := request.GetUserUUID()
	userOrganization, err := api.app.Dao().GetUserOrganizationByOrgIdAndUserUUID(orgId, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if userOrganization != nil && userOrganization.Role == "owner" {
		return models.NewErrorResponse(http.StatusForbidden, "Owner can't leave organization")
	}
	err = api.app.Dao().DeleteUserOrganizationFromEmailAndOrgId(email, orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, nil, "User left organization")
	return response
}

// RemoveOrg godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Remove user from organization.
//	@Description	Remove user from organization.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/remove [post]
//	@Param			orgId	path	string					true	"Organization ID"
//	@Param			data	body	models.UserEmailRequest	true	"User to remove"
func (api *Api) RemoveOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	var emailData string
	if email == nil {
		emailData = ""
	} else {
		emailData = email.(string)
	}
	if emailData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is required")
	}
	if addr, ok := coreservice.ValidateMailAddress(emailData); ok {
		emailData = addr
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is invalid")
	}
	orgId := request.GetOrgId()
	user, err := api.app.Dao().GetUserByEmail(emailData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	var response *models.Response
	if user == nil {
		response = models.NewErrorResponse(http.StatusNotFound, "User to remove not found")
		return response
	}
	userUUID := request.GetUserUUID()
	userOrganization, err := api.app.Dao().GetUserOrganizationByOrgIdAndUserUUID(orgId, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if userOrganization == nil || userOrganization.Role != "owner" {
		return models.NewErrorResponse(http.StatusForbidden, "You are not authorized to remove users from this organization")
	}
	userOrganization, err = api.app.Dao().GetUserOrganizationByOrgIdAndUserUUID(orgId, user.UUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if userOrganization == nil {
		return models.NewErrorResponse(http.StatusNotFound, "User not member of organization")
	}
	if userOrganization.Role == "owner" {
		return models.NewErrorResponse(http.StatusForbidden, "Owner can't be removed from organization")
	}
	err = api.app.Dao().DeleteUserOrganizationFromEmailAndOrgId(emailData, orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response = models.NewDataResponse(http.StatusOK, nil, "User removed from organization")
	return response
}

var GetOrgsForUser ServiceFunc = (*Api).GetOrgsForUser
var AddUsersToOrg ServiceFunc = (*Api).AddUsersToOrg
var UpdateUserRole ServiceFunc = (*Api).UpdateUserRole
var JoinOrg ServiceFunc = (*Api).JoinOrg
var LeaveOrg ServiceFunc = (*Api).LeaveOrg
var RemoveOrg ServiceFunc = (*Api).RemoveOrg
