package main

import (
	_ "github.com/PureML-Inc/PureML/server/docs"
	"github.com/PureML-Inc/PureML/server/handler"
	"github.com/PureML-Inc/PureML/server/middlewares"
	"github.com/PureML-Inc/PureML/server/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title PureML API Documentation
// @version 1.0

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email contact@pureml.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//Health API
	e.GET("/health", handler.DefaultHandler(service.HealthCheck))
	//Swagger API
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//Org APIs
	group := e.Group("/org")
	group.GET("/all", handler.DefaultHandler(service.GetAllAdminOrgs), middlewares.AuthenticateJWT)
	group.GET("/id/:orgId", handler.DefaultHandler(service.GetOrgByID), middlewares.AuthenticateJWT, middlewares.ValidateOrg)
	group.GET("/", handler.DefaultHandler(service.GetOrgsForUser), middlewares.AuthenticateJWT, middlewares.ValidateOrg)
	group.POST("/create", handler.DefaultHandler(service.CreateOrg), middlewares.AuthenticateJWT)
	group.POST("/:orgId/update", handler.DefaultHandler(service.UpdateOrg), middlewares.AuthenticateJWT, middlewares.ValidateOrg)
	group.POST("/:orgId/add", handler.DefaultHandler(service.UpdateOrg), middlewares.AuthenticateJWT, middlewares.ValidateOrg)
	group.POST("/join", handler.DefaultHandler(service.JoinOrg), middlewares.AuthenticateJWT)
	group.POST("/:orgId/remove", handler.DefaultHandler(service.RemoveOrg), middlewares.AuthenticateJWT, middlewares.ValidateOrg)
	group.POST("/:orgId/leave", handler.DefaultHandler(service.LeaveOrg), middlewares.AuthenticateJWT, middlewares.ValidateOrg)

	//Project APIs
	// group := e.Group("")

	//User APIs
	// group = e.Group("user")

	e.Logger.Fatal(e.Start("localhost:8080"))

}
