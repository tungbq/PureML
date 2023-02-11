package service

import (
	"fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
)

type ServiceFunc func (*Api, *models.Request) *models.Response

func (api *Api) DefaultHandler(f ServiceFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		request := extractRequest(context)
		response := f(api, request)
		responseWriter := context.Response().Writer
		if response.Error != nil {
			populateErrorResponse(context, response, responseWriter)
		} else {
			populateSuccessResponse(context, response, responseWriter)
		}
		return nil
	}
}

func populateSuccessResponse(context echo.Context, response *models.Response, responseWriter http.ResponseWriter) {
	context.Response().WriteHeader(response.StatusCode)
	_, err := responseWriter.Write(convertToBytes(response.Body))
	if err != nil {
		panic(fmt.Sprintf("Error writing response: %v \n", err.Error()))
	}
}

func populateErrorResponse(context echo.Context, response *models.Response, responseWriter http.ResponseWriter) {
	context.Response().WriteHeader(http.StatusInternalServerError)
	_, err := responseWriter.Write(convertToBytes(map[string]interface{}{
		"error": "Internal server error - " + response.Error.Error(),
	}))
	if err != nil {
		panic(fmt.Sprintf("Error writing response: %v \n", err.Error()))
	}
}
