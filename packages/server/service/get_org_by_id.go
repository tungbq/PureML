package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetOrgByID(request *models.Request) *models.Response {
	var response *models.Response
	orgId := request.PathParams["orgId"]
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
