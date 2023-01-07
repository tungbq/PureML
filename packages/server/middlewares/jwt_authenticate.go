package middlewares

import (
	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthenticateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		if !user.Valid {
			return echo.ErrUnauthorized
		}
		claims := user.Claims.(jwt.MapClaims)
		userHandle := claims["user_handle"].(string)
		userEmail := claims["user_email"].(string)
		c.Set("user_handle", userHandle)
		c.Set("user_email", userEmail)
		return next(c)
	}
}

// token := jwt.New(jwt.SigningMethodEdDSA)
// claims := token.Claims.(jwt.MapClaims)
// claims["exp"] = time.Now().Add(10 * time.Minute)
// claims["authorized"] = true
// claims["user"] = "username"
// tokenString, err := token.SignedString(sampleSecretKey)
// if err != nil {
//     return "", err
//  }

//  return tokenString, nil
