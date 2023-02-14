package middlewares

import (
	_ "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func ValidateDatasetBranch(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			datasetBranchName := context.Param("branchName")
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
			if datasetBranchName == "" {
				context.Response().WriteHeader(http.StatusBadRequest)
				_, err = context.Response().Writer.Write([]byte("Branch name required"))
				if err != nil {
					return err
				}
				return nil
			}
			branch, err := app.Dao().GetDatasetBranchByName(orgUUID, datasetName, datasetBranchName)
			if err != nil {
				context.Response().WriteHeader(http.StatusInternalServerError)
				_, err = context.Response().Writer.Write([]byte(err.Error()))
				if err != nil {
					return err
				}
				return nil
			}
			if branch == nil {
				context.Response().WriteHeader(http.StatusNotFound)
				_, err = context.Response().Writer.Write([]byte("Branch not found"))
				if err != nil {
					return err
				}
				return nil
			}
			context.Set(ContextDatasetBranchKey, &models.DatasetBranchNameResponse{
				Name: branch.Name,
				UUID: branch.UUID,
			})
			return next(context)
		}
	}
}
