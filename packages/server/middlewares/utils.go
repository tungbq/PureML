package middlewares

import "github.com/labstack/echo/v5"

func extractRequestHeader(headerName string, context echo.Context) string {
	headerValue := ""
	headerValues := context.Request().Header[AuthHeaderName]
	if len(headerValues) != 0 {
		headerValue = headerValues[0]
	}
	return headerValue
}
