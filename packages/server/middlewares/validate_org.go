package middlewares

import (
	"net/http"

	ds "github.com/PureML-Inc/PureML/server/daos"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func ValidateOrg(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		orgId := context.Param("orgId")
		if orgId == "" {
			context.Response().WriteHeader(http.StatusBadRequest)
			context.Response().Writer.Write([]byte("Organization Id required"))
			return nil
		}
		org, err := ds.GetOrgById(uuid.Must(uuid.FromString(orgId)))
		if err != nil {
			context.Response().WriteHeader(http.StatusInternalServerError)
			context.Response().Writer.Write([]byte(err.Error()))
			return nil
		}
		if org == nil {
			context.Response().WriteHeader(http.StatusNotFound)
			context.Response().Writer.Write([]byte("Organization not found"))
			return nil
		}
		context.Set("Org", &models.OrganizationHandleResponse{
			Name:   org.Name,
			UUID:   org.UUID,
			Handle: org.Handle,
			Avatar: org.Avatar,
		})
		return next(context)
	}
}
