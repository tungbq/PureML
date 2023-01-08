package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetOrgByID(request *models.Request) *models.Response {
	response := &models.Response{}
	orgId := request.PathParams["orgId"]
	organization, err := datastore.GetOrgById(orgId)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	if organization == nil {
		response.StatusCode = http.StatusNotFound
		response.Body.Status = response.StatusCode
		response.Body.Message = "Organization not found"
		response.Body.Data = nil
	} else {
		response.StatusCode = http.StatusOK
		response.Body.Status = response.StatusCode
		response.Body.Message = "Organization Details"
		response.Body.Data = []models.Organization{*organization}
	}
	return response
}
