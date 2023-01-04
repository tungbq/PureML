package main

import (
	"github.com/PriyavKaneria/PureML/service/handlers"
	"github.com/PriyavKaneria/PureML/service/middlewares"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//Org APIs
	group := e.Group("/org")
	group.GET("/all", handlers.GetAllAdminOrgs, middlewares.AuthenticateJWT)

	//Project APIs
	// group := e.Group("")

	//User APIs
	// group = e.Group("user")

	e.Logger.Fatal(e.Start(":8080"))

}
