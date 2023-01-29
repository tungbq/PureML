package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

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
func GetOrgByID(request *models.Request) *models.Response {
	var response *models.Response
	orgId := uuid.Must(uuid.FromString(request.PathParams["orgId"]))
	organization, err := datastore.GetOrgById(orgId)
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
