package middlewares

import (
	"net/http"

	ds "github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
)

func ValidateModelBranch(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		branchName := context.Param("branchName")
		modelName := context.Param("modelName")
		if branchName == "" {
			context.Response().WriteHeader(http.StatusBadRequest)
			context.Response().Writer.Write([]byte("Branch name required"))
			return nil
		}
		branch, err := ds.GetBranchByName(modelName, branchName)
		if err != nil {
			context.Response().WriteHeader(http.StatusInternalServerError)
			context.Response().Writer.Write([]byte(err.Error()))
			return nil
		}
		if branch == nil {
			context.Response().WriteHeader(http.StatusNotFound)
			context.Response().Writer.Write([]byte("Branch not found"))
			return nil
		}
		context.Set("ModelBranch", &models.ModelBranchNameResponse{
			Name: branch.Name,
			UUID: branch.UUID,
		})
		return next(context)
	}
}
