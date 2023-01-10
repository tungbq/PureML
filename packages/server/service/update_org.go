package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// UpdateOrg godoc
// @Security ApiKeyAuth
// @Summary Update organization details.
// @Description Update organization details by ID.
// @Tags Organization
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/update [post]
// @Param orgId path string true "Organization ID"
// @Param org body models.CreateOrUpdateOrgRequest true "Organization details"
func UpdateOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := uuid.Must(uuid.FromString(request.PathParams["orgId"]))
	orgName := request.GetParsedBodyAttribute("name").(string)
	orgDesc := request.GetParsedBodyAttribute("description").(string)
	orgAvatar := request.GetParsedBodyAttribute("avatar").(string)
	email := request.User.Email
	UserOrganization, err := datastore.GetUserOrganizationByOrgIdAndEmail(orgId, email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	var response *models.Response
	if UserOrganization.Role != "owner" {
		response = models.NewErrorResponse(http.StatusForbidden, "You are not authorized to update this organization")
		return response
	}
	updatedOrg, err := datastore.UpdateOrg(orgId, orgName, orgDesc, orgAvatar)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response = models.NewDataResponse(http.StatusOK, []models.OrganizationResponse{*updatedOrg}, "Organization updated")
	return response

}
