package middlewares

import (
	"cqrses/controllers"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// IsAdmin is the middleware function.
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		admin := claims["admin"].(bool)

		if !admin {
			return c.JSON(
				http.StatusUnauthorized,
				controllers.SetResponse(
					http.StatusUnauthorized,
					"Authorization error",
					"User is not an administrator",
				),
			)
		}

		return next(c)
	}
}
