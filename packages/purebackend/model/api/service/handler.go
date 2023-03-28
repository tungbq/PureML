package service

import (
	"fmt"
	"net/http"
	"strings"

	coreservice "github.com/PureMLHQ/PureML/packages/purebackend/core/apis/service"
	"github.com/PureMLHQ/PureML/packages/purebackend/core/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type ServiceFunc func(*Api, *models.Request) *models.Response

func (api *Api) DefaultHandler(f ServiceFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		request := coreservice.ExtractRequest(context)
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
	_, err := responseWriter.Write(coreservice.ConvertToBytes(response.Body))
	if err != nil {
		panic(fmt.Sprintf("Error writing response: %v \n", err.Error()))
	}
}

func populateErrorResponse(context echo.Context, response *models.Response, responseWriter http.ResponseWriter) {
	context.Response().WriteHeader(http.StatusInternalServerError)
	_, err := responseWriter.Write(coreservice.ConvertToBytes(map[string]interface{}{
		"error": "Internal server error - " + response.Error.Error(),
	}))
	if err != nil {
		panic(fmt.Sprintf("Error writing response: %v \n", err.Error()))
	}
}

func (api *Api) ValidateSourceTypeAndGetPublicURL(modelSourceType string, orgId uuid.UUID) (string, *models.Response) {
	var sourceTypePublicURL string
	var err error
	modelSourceType = strings.ToUpper(modelSourceType)
	if modelSourceType == "PUREML-STORAGE" {
		sourceTypePublicURL, err = api.app.Dao().GetSourcePublicURL(uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), "R2")
	} else {
		sourceTypePublicURL, err = api.app.Dao().GetSourcePublicURL(orgId, modelSourceType)
	}
	if sourceTypePublicURL == "" || err != nil {
		return "", models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Source %s not connected properly to organization", modelSourceType))
	}
	return sourceTypePublicURL, nil
}
