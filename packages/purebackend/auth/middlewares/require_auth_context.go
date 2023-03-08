package middlewares

import (
	"net/http"

	"github.com/PureMLHQ/PureML/packages/purebackend/user_org/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func RequireAuthContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		userContext := context.Get(ContextAuthKey)
		if userContext == nil {
			context.Response().WriteHeader(http.StatusUnauthorized)
			_, err := context.Response().Writer.Write([]byte("Authentication token required"))
			if err != nil {
				return err
			}
			return nil
		} else {
			userContext := userContext.(*models.UserClaims)
			if userContext.UUID == uuid.Nil {
				context.Response().WriteHeader(http.StatusNotFound)
				_, err := context.Response().Writer.Write([]byte("User not found"))
				if err != nil {
					return err
				}
				return nil
			}
		}
		return next(context)
	}
}
