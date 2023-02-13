package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/middlewares"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// BindOrgApi registers the admin api endpoints and the corresponding handlers.
func BindOrgApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	rg.GET("/org/handle/:orgHandle", api.DefaultHandler(GetOrgByHandle))
	rg.GET("/org/:orgId/public/model", api.DefaultHandler(GetOrgAllPublicModels), middlewares.ValidateOrg(api.app))
	rg.GET("/org/:orgId/public/dataset", api.DefaultHandler(GetOrgAllPublicDatasets), middlewares.ValidateOrg(api.app))
	orgGroup := rg.Group("/org", middlewares.RequireAuthContext)
	orgGroup.GET("/id/:orgId", api.DefaultHandler(GetOrgByID), middlewares.ValidateOrg(api.app))
	orgGroup.POST("/create", api.DefaultHandler(CreateOrg))
	orgGroup.POST("/:orgId/update", api.DefaultHandler(UpdateOrg), middlewares.ValidateOrg(api.app))
}

// GetOrgByHandle godoc
//
//	@Summary		Get organization details by handle.
//	@Description	Get organization details by handle.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/handle/{orgHandle} [get]
//	@Param			orgHandle	path	string	true	"Organization Handle"
func (api *Api) GetOrgByHandle(request *models.Request) *models.Response {
	var response *models.Response
	orgHandle := request.PathParams["orgHandle"]
	organization, err := api.app.Dao().GetOrgByHandle(orgHandle)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if organization == nil {
		response = models.NewErrorResponse(http.StatusNotFound, "Organization not found")
	} else {
		response = models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*organization}, "Organization details")
	}
	return response
}

// GetOrgByID godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get organization details by ID.
//	@Description	Get organization details by ID.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/id/{orgId} [get]
//	@Param			orgId	path	string	true	"Organization ID"
func (api *Api) GetOrgByID(request *models.Request) *models.Response {
	var response *models.Response
	orgId := uuid.Must(uuid.FromString(request.PathParams["orgId"]))
	organization, err := api.app.Dao().GetOrgById(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if organization == nil {
		response = models.NewErrorResponse(http.StatusNotFound, "Organization not found")
	} else {
		response = models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*organization}, "Organization details")
	}
	return response
}

// GetOrgAllPublicModels godoc
//
//	@Summary		Get all public models of an organization.
//	@Description	Get all public models of an organization.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/public/model [get]
//	@Param			orgId	path	string	true	"Organization ID"
func (api *Api) GetOrgAllPublicModels(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	modelsdb, err := api.app.Dao().GetOrgAllPublicModels(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, modelsdb, "Public models of Organization")
}

// GetOrgAllPublicDatasets godoc
//
//	@Summary		Get all public datasets of an organization.
//	@Description	Get all public datasets of an organization.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/public/dataset [get]
//	@Param			orgId	path	string	true	"Organization ID"
func (api *Api) GetOrgAllPublicDatasets(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	datasetsdb, err := api.app.Dao().GetOrgAllPublicDatasets(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, datasetsdb, "Public datasets of Organization")
}

// CreateOrg godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create an organization.
//	@Description	Create an organization and add the user as the owner.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/create [post]
//	@Param			org	body	models.CreateOrgRequest	true	"Organization details"
func (api *Api) CreateOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	// orgName := request.GetParsedBodyAttribute("name").(string)
	orgDesc := request.GetParsedBodyAttribute("description")
	var orgDescData string
	if orgDesc == nil {
		orgDescData = ""
	} else {
		orgDescData = orgDesc.(string)
	}
	orgHandle := request.GetParsedBodyAttribute("handle")
	var orgHandleData string
	if orgHandle == nil {
		orgHandleData = ""
	} else {
		orgHandleData = orgHandle.(string)
	}
	if orgHandleData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Organization handle is required")
	}
	orgName := orgHandleData
	email := request.User.Email
	org, err := api.app.Dao().CreateOrgFromEmail(email, orgName, orgDescData, orgHandleData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*org}, "Organization created")
	return response
}

// UpdateOrg godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update organization details.
//	@Description	Update organization details by ID.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/update [post]
//	@Param			orgId	path	string					true	"Organization ID"
//	@Param			org		body	models.UpdateOrgRequest	true	"Organization details"
func (api *Api) UpdateOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := uuid.Must(uuid.FromString(request.PathParams["orgId"]))
	userUUID := request.GetUserUUID()
	UserOrganization, err := api.app.Dao().GetUserOrganizationByOrgIdAndUserUUID(orgId, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if UserOrganization == nil || UserOrganization.Role != "owner" {
		return models.NewErrorResponse(http.StatusForbidden, "You are not authorized to update this organization")
	}
	orgName := request.GetParsedBodyAttribute("name")
	orgDesc := request.GetParsedBodyAttribute("description")
	orgAvatar := request.GetParsedBodyAttribute("avatar")
	updatedAttributes := map[string]interface{}{}
	if orgName != nil {
		updatedAttributes["name"] = orgName.(string)
		if orgName == "" {
			return models.NewErrorResponse(http.StatusBadRequest, "Name cannot be empty")
		}
	}
	if orgAvatar != nil {
		updatedAttributes["avatar"] = orgAvatar.(string)
	}
	if orgDesc != nil {
		updatedAttributes["desc"] = orgDesc.(string)
	}
	updatedOrg, err := api.app.Dao().UpdateOrg(orgId, updatedAttributes)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*updatedOrg}, "Organization updated")
}

var GetOrgByHandle ServiceFunc = (*Api).GetOrgByHandle
var GetOrgByID ServiceFunc = (*Api).GetOrgByID
var GetOrgAllPublicModels ServiceFunc = (*Api).GetOrgAllPublicModels
var GetOrgAllPublicDatasets ServiceFunc = (*Api).GetOrgAllPublicDatasets
var CreateOrg ServiceFunc = (*Api).CreateOrg
var UpdateOrg ServiceFunc = (*Api).UpdateOrg
