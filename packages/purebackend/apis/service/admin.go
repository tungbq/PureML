package service

import (
	// "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/packages/purebackend/config"
	"github.com/PureML-Inc/PureML/packages/purebackend/core"
	"github.com/PureML-Inc/PureML/packages/purebackend/middlewares"
	"github.com/PureML-Inc/PureML/packages/purebackend/models"
	"github.com/labstack/echo/v4"
)

// BindAdminApi registers the admin api endpoints and the corresponding handlers.
func BindAdminApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	orgGroup := rg.Group("/org", middlewares.RequireAuthContext)
	orgGroup.GET("/all", api.DefaultHandler(GetAllAdminOrgs))
}

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
func (api *Api) GetAllAdminOrgs(request *models.Request) *models.Response {
	var response *models.Response
	if request.User == nil {
		response = models.NewErrorResponse(http.StatusUnauthorized, "Unauthorized")
		return response
	}
	if config.HasAdminAccess(request.User.Email) {
		allOrgs, err := api.app.Dao().GetAllAdminOrgs()
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

var GetAllAdminOrgs ServiceFunc = (*Api).GetAllAdminOrgs
