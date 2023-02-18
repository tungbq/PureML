package middlewares

import (
	"net/http"

	"github.com/PureML-Inc/PureML/packages/purebackend/core"
	"github.com/PureML-Inc/PureML/packages/purebackend/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func ValidateModel(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			modelName := context.Param("modelName")
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
			if modelName == "" {
				context.Response().WriteHeader(http.StatusBadRequest)
				_, err = context.Response().Writer.Write([]byte("Model name required"))
				if err != nil {
					return err
				}
				return nil
			}
			model, err := app.Dao().GetModelByName(orgUUID, modelName)
			if err != nil {
				context.Response().WriteHeader(http.StatusInternalServerError)
				_, err = context.Response().Writer.Write([]byte(err.Error()))
				if err != nil {
					return err
				}
				return nil
			}
			if model == nil {
				context.Response().WriteHeader(http.StatusNotFound)
				_, err = context.Response().Writer.Write([]byte("Model not found"))
				if err != nil {
					return err
				}
				return nil
			}
			context.Set(ContextModelKey, &models.ModelNameResponse{
				Name: model.Name,
				UUID: model.UUID,
			})
			return next(context)
		}
	}
}
