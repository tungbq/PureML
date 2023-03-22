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

func (api *Api) ValidateAndGetOrCreateSourceType(datasetSourceType string, orgId uuid.UUID) (uuid.UUID, *models.Response) {
	var sourceTypeUUID uuid.UUID
	var err error
	datasetSourceType = strings.ToUpper(datasetSourceType)
	if datasetSourceType == "PUREML-STORAGE" {
		sourceTypeUUID, err = api.app.Dao().GetSourceTypeByName(uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), datasetSourceType)
		if sourceTypeUUID == uuid.Nil || err != nil {
			return uuid.Nil, models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Source %s not connected properly to organization", datasetSourceType))
		}
	} else {
		sourceTypeUUID, err = api.app.Dao().GetSourceTypeByName(orgId, datasetSourceType)
		if err != nil {
			return uuid.Nil, models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Source %s not connected properly to organization", datasetSourceType))
		}
		if sourceTypeUUID == uuid.Nil {
			if datasetSourceType == "S3" && api.app.Settings().S3.Enabled {
				publicUrl := fmt.Sprintf("https://%s.%s", api.app.Settings().S3.Bucket, api.app.Settings().S3.Endpoint)
				sourceType, err := api.app.Dao().CreateS3Source(orgId, publicUrl)
				if err != nil {
					return uuid.Nil, models.NewServerErrorResponse(err)
				}
				sourceTypeUUID = sourceType.UUID
			} else if datasetSourceType == "R2" && api.app.Settings().R2.Enabled {
				publicUrl := fmt.Sprintf("https://%s/%s", api.app.Settings().R2.Endpoint, api.app.Settings().R2.Bucket)
				sourceType, err := api.app.Dao().CreateR2Source(orgId, publicUrl)
				if err != nil {
					return uuid.Nil, models.NewServerErrorResponse(err)
				}
				sourceTypeUUID = sourceType.UUID
			} else if datasetSourceType == "LOCAL" {
				sourceType, err := api.app.Dao().CreateLocalSource(orgId)
				if err != nil {
					return uuid.Nil, models.NewServerErrorResponse(err)
				}
				sourceTypeUUID = sourceType.UUID
			} else {
				return uuid.Nil, models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("%s source not enabled", datasetSourceType))
			}
		}
	}
	return sourceTypeUUID, nil
}
