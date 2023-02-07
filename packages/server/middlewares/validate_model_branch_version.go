package middlewares

import (
	_ "fmt"
	"net/http"

	ds "github.com/PureML-Inc/PureML/server/daos"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
)

func ValidateModelBranchVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		modelBranchVersion := context.Param("version")
		modelBranchUUID := context.Get("ModelBranch").(*models.ModelBranchNameResponse).UUID
		if modelBranchVersion == "" {
			context.Response().WriteHeader(http.StatusBadRequest)
			context.Response().Writer.Write([]byte("Version required"))
			return nil
		}
		version, err := ds.GetModelBranchVersion(modelBranchUUID, modelBranchVersion)
		if err != nil {
			context.Response().WriteHeader(http.StatusInternalServerError)
			context.Response().Writer.Write([]byte(err.Error()))
			return nil
		}
		if version == nil {
			context.Response().WriteHeader(http.StatusNotFound)
			context.Response().Writer.Write([]byte("Model Branch Version not found"))
			return nil
		}
		context.Set("ModelBranchVersion", &models.ModelBranchVersionNameResponse{
			UUID:    version.UUID,
			Version: version.Version,
		})
		return next(context)
	}
}
