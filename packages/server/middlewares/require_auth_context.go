package middlewares

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func RequireAuthContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		userContext := context.Get(ContextAuthKey)
		if userContext == nil {
			context.Response().WriteHeader(http.StatusUnauthorized)
			context.Response().Writer.Write([]byte("Authentication token required"))
			return nil
		} else {
			userContext := userContext.(*models.UserClaims)
			if userContext.UUID == uuid.Nil {
				context.Response().WriteHeader(http.StatusNotFound)
				context.Response().Writer.Write([]byte("User not found"))
				return nil
			}
		}
		return next(context)
	}
}
