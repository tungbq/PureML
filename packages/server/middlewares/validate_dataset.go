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
				_, err = context.Response().Writer.Write([]byte("Invalid UUID format"))
				if err != nil {
					return err
				}
				return nil
			}
			if datasetName == "" {
				context.Response().WriteHeader(http.StatusBadRequest)
				_, err = context.Response().Writer.Write([]byte("Dataset name required"))
				if err != nil {
					return err
				}
				return nil
			}
			dataset, err := app.Dao().GetDatasetByName(orgUUID, datasetName)
			if err != nil {
				context.Response().WriteHeader(http.StatusInternalServerError)
				_, err = context.Response().Writer.Write([]byte(err.Error()))
				if err != nil {
					return err
				}
				return nil
			}
			if dataset == nil {
				context.Response().WriteHeader(http.StatusNotFound)
				_, err = context.Response().Writer.Write([]byte("Dataset not found"))
				if err != nil {
					return err
				}
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
