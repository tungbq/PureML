package middlewares

import "github.com/labstack/echo/v4"

func ExtractRequestHeader(headerName string, context echo.Context) string {
	headerValue := ""
	headerValues := context.Request().Header[headerName]
	if len(headerValues) != 0 {
		headerValue = headerValues[0]
	}
	return headerValue
}
