package middlewares

import (
	"net/http"

	ds "github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func ValidateModel(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		modelName := context.Param("modelName")
		orgId := uuid.Must(uuid.FromString(context.Param("orgId")))
		if modelName == "" {
			context.Response().WriteHeader(http.StatusBadRequest)
			context.Response().Writer.Write([]byte("Model name required"))
			return nil
		}
		model, err := ds.GetModelByName(orgId, modelName)
		if err != nil {
			context.Response().WriteHeader(http.StatusInternalServerError)
			context.Response().Writer.Write([]byte(err.Error()))
			return nil
		}
		if model == nil {
			context.Response().WriteHeader(http.StatusNotFound)
			context.Response().Writer.Write([]byte("Model not found"))
			return nil
		}
		context.Set("Model", &models.ModelNameResponse{
			Name: model.Name,
			UUID: model.UUID,
		})
		return next(context)
	}
}
