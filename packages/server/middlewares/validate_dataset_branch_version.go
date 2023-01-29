package middlewares

import (
	_ "fmt"
	"net/http"

	ds "github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
)

func ValidateDatasetBranchVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		datasetBranchVersion := context.Param("version")
		datasetBranchUUID := context.Get("DatasetBranch").(*models.DatasetBranchNameResponse).UUID
		if datasetBranchVersion == "" {
			context.Response().WriteHeader(http.StatusBadRequest)
			context.Response().Writer.Write([]byte("Version required"))
			return nil
		}
		version, err := ds.GetDatasetBranchVersion(datasetBranchUUID, datasetBranchVersion)
		if err != nil {
			context.Response().WriteHeader(http.StatusInternalServerError)
			context.Response().Writer.Write([]byte(err.Error()))
			return nil
		}
		if version == nil {
			context.Response().WriteHeader(http.StatusNotFound)
			context.Response().Writer.Write([]byte("Dataset Branch Version not found"))
			return nil
		}
		context.Set("DatasetBranchVersion", &models.DatasetBranchVersionNameResponse{
			UUID:    version.UUID,
			Version: version.Version,
		})
		return next(context)
	}
}
