package main

import (
	"github.com/PriyavKaneria/PureML/service/handler"
	"github.com/PriyavKaneria/PureML/service/middlewares"
	"github.com/PriyavKaneria/PureML/service/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

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

	e.Logger.Fatal(e.Start(":8080"))

}
