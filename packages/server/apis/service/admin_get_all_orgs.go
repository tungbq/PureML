package service

import (
	// "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetAllAdminOrgs godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all organizations and their details.
//	@Description	Get all organizations and their details. Only accessible by admins.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/all [get]
func GetAllAdminOrgs(request *models.Request) *models.Response {
	var response *models.Response
	if request.User == nil {
		response = models.NewErrorResponse(http.StatusUnauthorized, "Unauthorized")
		return response
	}
	if config.HasAdminAccess(request.User.Email) {
		allOrgs, err := datastore.GetAllAdminOrgs()
		if err != nil {
			return models.NewServerErrorResponse(err)
		} else {
			response = models.NewDataResponse(http.StatusOK, allOrgs, "All organizations")
		}
	} else {
		response = models.NewErrorResponse(http.StatusForbidden, "Forbidden")
	}
	return response
}
