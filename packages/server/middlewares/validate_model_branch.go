package middlewares

import (
	_ "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func ValidateModelBranch(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			branchName := context.Param("branchName")
			modelName := context.Param("modelName")
			orgId := context.Param("orgId")
			orgUUID, err := uuid.FromString(orgId)
			if err != nil {
				context.Response().WriteHeader(http.StatusBadRequest)
				context.Response().Writer.Write([]byte("Invalid UUID format"))
				return nil
			}
			if branchName == "" {
				context.Response().WriteHeader(http.StatusBadRequest)
				context.Response().Writer.Write([]byte("Branch name required"))
				return nil
			}
			branch, err := app.Dao().GetModelBranchByName(orgUUID, modelName, branchName)
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
			context.Set(ContextModelBranchKey, &models.ModelBranchNameResponse{
				Name: branch.Name,
				UUID: branch.UUID,
			})
			return next(context)
		}
	}
}
