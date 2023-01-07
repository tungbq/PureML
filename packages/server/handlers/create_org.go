package handlers

import (
	"github.com/PriyavKaneria/PureML/service/service"
	"github.com/labstack/echo/v4"
)

// CreateOrganization godoc
// @Summary Create a test organization.
// @Description Create a test organization.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/create [post]
func CreateOrganization(context echo.Context) error {
	request := extractRequest(context)
	response := service.CreateOrganization(request)
	var err error
	context.Response().WriteHeader(response.StatusCode)
	responseWriter := context.Response().Writer
	_, err = responseWriter.Write(convertToBytes(response.Body))
	return err
}
