package handlers

import (
	"github.com/PriyavKaneria/PureML/service/service"
	"github.com/labstack/echo/v4"
)

func GetAllAdminOrgs(context echo.Context) error {
	request := extractRequest(context)
	response := service.GetAllAdminOrgs(request)
	var err error
	if response.Error != nil {
		err = response.Error
	} else {
		context.Response().WriteHeader(response.StatusCode)
		responseWriter := context.Response().Writer
		_, err = responseWriter.Write(convertToBytes(response.Body))
	}
	return err
}
