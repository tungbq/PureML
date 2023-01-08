package service

import (
	"fmt"
	// "github.com/PriyavKaneria/PureML/service/config"
	"github.com/PriyavKaneria/PureML/service/datastore"
	"github.com/PriyavKaneria/PureML/service/models"
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetAllAdminOrgs(request *models.Request) *models.Response {
	response := &models.Response{}
	if config.HasAdminAccess(request.User.MailId) {
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
		response.Body = "Forbidden"
	}
	return response
}

func CreateOrganization(request *models.Request) *models.Response {
	response := &models.Response{}
	org := models.Organization{
		Name:         "TestOrg",
		Handle:       "testorg",
		Avatar:       "",
		Description:  "Test org",
		APITokenHash: "",
		JoinCode:     "testjoincode",
	}
	err := datastore.CreateOrganization(org)
	if err != nil {
		fmt.Println(err)
		response.Error = err
		response.StatusCode = http.StatusInternalServerError
		response.Body.Status = response.StatusCode
		response.Body.Message = fmt.Sprintf("Internal server error - %s", err.Error())
		response.Body.Data = nil
	} else {
		response.StatusCode = http.StatusOK
		response.Body.Status = response.StatusCode
		response.Body.Message = "Organization created"
		response.Body.Data = nil
	}

	return response
}
