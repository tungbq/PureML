package middlewares

import (
	_ "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/packages/purebackend/core"
	"github.com/PureML-Inc/PureML/packages/purebackend/models"
	"github.com/labstack/echo/v4"
)

func ValidateDatasetBranchVersion(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			datasetBranchVersion := context.Param("version")
			datasetBranchUUID := context.Get("DatasetBranch").(*models.DatasetBranchNameResponse).UUID
			if datasetBranchVersion == "" {
				context.Response().WriteHeader(http.StatusBadRequest)
				_, err := context.Response().Writer.Write([]byte("Version required"))
				if err != nil {
					return err
				}
				return nil
			}
			version, err := app.Dao().GetDatasetBranchVersion(datasetBranchUUID, datasetBranchVersion)
			if err != nil {
				context.Response().WriteHeader(http.StatusInternalServerError)
				_, err = context.Response().Writer.Write([]byte(err.Error()))
				if err != nil {
					return err
				}
				return nil
			}
			if version == nil {
				context.Response().WriteHeader(http.StatusNotFound)
				_, err = context.Response().Writer.Write([]byte("Dataset Branch Version not found"))
				if err != nil {
					return err
				}
				return nil
			}
			context.Set(ContextDatasetBranchVersionKey, &models.DatasetBranchVersionNameResponse{
				UUID:    version.UUID,
				Version: version.Version,
			})
			return next(context)
		}
	}
}
