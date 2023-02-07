package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

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
		response = models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*organization}, "Organization Details")
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
		response = models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*organization}, "Organization Details")
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
//	@Param			org	body	models.CreateOrUpdateOrgRequest	true	"Organization details"
func (api *Api) CreateOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	// orgName := request.GetParsedBodyAttribute("name").(string)
	orgDesc := request.GetParsedBodyAttribute("description").(string)
	orgHandle := request.GetParsedBodyAttribute("handle").(string)
	orgName := orgHandle
	email := request.User.Email
	org, err := api.app.Dao().CreateOrgFromEmail(email, orgName, orgDesc, orgHandle)
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
//	@Param			orgId	path	string							true	"Organization ID"
//	@Param			org		body	models.CreateOrUpdateOrgRequest	true	"Organization details"
func (api *Api) UpdateOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := uuid.Must(uuid.FromString(request.PathParams["orgId"]))
	orgName := request.GetParsedBodyAttribute("name").(string)
	orgDesc := request.GetParsedBodyAttribute("description").(string)
	orgAvatar := request.GetParsedBodyAttribute("avatar").(string)
	email := request.User.Email
	UserOrganization, err := api.app.Dao().GetUserOrganizationByOrgIdAndEmail(orgId, email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	var response *models.Response
	if UserOrganization.Role != "owner" {
		response = models.NewErrorResponse(http.StatusForbidden, "You are not authorized to update this organization")
		return response
	}
	updatedOrg, err := api.app.Dao().UpdateOrg(orgId, orgName, orgDesc, orgAvatar)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response = models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*updatedOrg}, "Organization updated")
	return response

}
