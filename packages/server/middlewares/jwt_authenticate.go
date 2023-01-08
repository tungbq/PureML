package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

const AuthHeaderName = "authorization"

func AuthenticateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		authHeaderValue := extractRequestHeader(AuthHeaderName, context)
		if authHeaderValue == "" {
			context.Response().WriteHeader(http.StatusUnauthorized)
			context.Response().Writer.Write([]byte("Authentication Token Required"))
			return nil
		}
		authHeaderValue = strings.Split(authHeaderValue, " ")[1] //Splitting the bearer part??
		token, _ := jwt.Parse(authHeaderValue, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid token signing algorithm %v", t.Method.Alg())
			}
			return config.TokenSigningSecret(), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			context.Set("User", claims["user"]) //Todo to create User object
			next(context)
		} else {
			context.Response().WriteHeader(http.StatusForbidden)
			context.Response().Writer.Write([]byte("Invalid Authentication Token"))
		}
		return nil
	}
}
