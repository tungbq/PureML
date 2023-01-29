package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetOrgsForUser godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all user organizations.
//	@Description	Get all user organizations.
//	@Tags			Organization
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/ [get]
func GetOrgsForUser(request *models.Request) *models.Response {
	email := request.GetUserMail()
	UserOrganization, err := datastore.GetUserOrganizationsByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusAccepted, UserOrganization, "User Organizations")
	return response
}
