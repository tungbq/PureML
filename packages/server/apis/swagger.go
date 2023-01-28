// TEMPORARY SWAGGER HANDLER FOR echo-v4 and v5 COMPATIBILITY
package apis

import (
	v4 "github.com/labstack/echo/v4"
	v5 "github.com/labstack/echo/v5"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SwaggerHandler(next v5.Context) error {
	handler := echoSwagger.WrapHandler
	var v4Context v4.Context
	return handler(v4Context)
}
