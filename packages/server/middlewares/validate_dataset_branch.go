package middlewares

import (
	_ "fmt"
	"net/http"

	ds "github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func ValidateDatasetBranch(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		datasetBranchName := context.Param("branchName")
		datasetName := context.Param("datasetName")
		orgId := uuid.Must(uuid.FromString(context.Param("orgId")))
		if datasetBranchName == "" {
			context.Response().WriteHeader(http.StatusBadRequest)
			context.Response().Writer.Write([]byte("Branch name required"))
			return nil
		}
		branch, err := ds.GetDatasetBranchByName(orgId, datasetName, datasetBranchName)
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
		context.Set("DatasetBranch", &models.DatasetBranchNameResponse{
			Name: branch.Name,
			UUID: branch.UUID,
		})
		return next(context)
	}
}
