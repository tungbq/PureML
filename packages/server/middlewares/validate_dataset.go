package middlewares

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func ValidateDataset(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			datasetName := context.Param("datasetName")
			orgId := context.Param("orgId")
			orgUUID, err := uuid.FromString(orgId)
			if err != nil {
				context.Response().WriteHeader(http.StatusBadRequest)
				context.Response().Writer.Write([]byte("Invalid UUID format"))
				return nil
			}
			if datasetName == "" {
				context.Response().WriteHeader(http.StatusBadRequest)
				context.Response().Writer.Write([]byte("Dataset name required"))
				return nil
			}
			dataset, err := app.Dao().GetDatasetByName(orgUUID, datasetName)
			if err != nil {
				context.Response().WriteHeader(http.StatusInternalServerError)
				context.Response().Writer.Write([]byte(err.Error()))
				return nil
			}
			if dataset == nil {
				context.Response().WriteHeader(http.StatusNotFound)
				context.Response().Writer.Write([]byte("Dataset not found"))
				return nil
			}
			context.Set(ContextDatasetKey, &models.DatasetNameResponse{
				Name: dataset.Name,
				UUID: dataset.UUID,
			})
			return next(context)
		}
	}
}
