package service

import (
	// "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetAllAdminOrgs(request *models.Request) *models.Response {
	response := &models.Response{}
	if config.HasAdminAccess(request.User.Email) {
		allOrgs, err := datastore.GetAllAdminOrgs()
		if err != nil {
			return models.NewErrorResponse(err)
		} else {
			response.StatusCode = http.StatusOK
			response.Body.Status = response.StatusCode
			response.Body.Message = "All organizations"
			response.Body.Data = allOrgs
		}
	} else {
		response.StatusCode = http.StatusForbidden
		response.Body.Status = response.StatusCode
		response.Body.Message = "Forbidden"
		response.Body.Data = nil
	}
	return response
}
