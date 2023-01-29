package middlewares

import (
	"net/http"

	ds "github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func ValidateDataset(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		datasetName := context.Param("datasetName")
		orgId := uuid.Must(uuid.FromString(context.Param("orgId")))
		if datasetName == "" {
			context.Response().WriteHeader(http.StatusBadRequest)
			context.Response().Writer.Write([]byte("Dataset name required"))
			return nil
		}
		dataset, err := ds.GetDatasetByName(orgId, datasetName)
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
		context.Set("Dataset", &models.DatasetNameResponse{
			Name: dataset.Name,
			UUID: dataset.UUID,
		})
		return next(context)
	}
}
