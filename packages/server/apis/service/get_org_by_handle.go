package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
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
func GetOrgByHandle(request *models.Request) *models.Response {
	var response *models.Response
	orgHandle := request.PathParams["orgHandle"]
	organization, err := datastore.GetOrgByHandle(orgHandle)
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
