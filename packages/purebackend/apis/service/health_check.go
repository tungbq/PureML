package service

import (
	"net/http"

	"github.com/PuremlHQ/PureML/packages/purebackend/core"
	"github.com/PuremlHQ/PureML/packages/purebackend/models"
	"github.com/labstack/echo/v4"
)

// BindHealthApi registers the admin api endpoints and the corresponding handlers.
func BindHealthApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	rg.GET("/health", api.DefaultHandler(HealthCheck))
}

// HealthCheck godoc
//
//	@Summary		Show the status of server.
//	@Description	Get the status of server.
//	@Tags			Root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/health [get]
func (api *Api) HealthCheck(request *models.Request) *models.Response {
	return models.NewDataResponse(http.StatusOK, nil, "Server is up and running🚀🎉")
}

var HealthCheck ServiceFunc = (*Api).HealthCheck
