package middlewares

import (
	"net/http"
	"regexp"

	"cqrses/app/controllers"

	"github.com/labstack/echo/v4"
)

// ParamValidation validate the parameter of the request.
func ParamValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramKey := c.ParamNames()
		paramValue := c.ParamValues()

		r := regexp.MustCompile("^[0-9]+$")

		for k, v := range paramValue {
			if !r.MatchString(v) {
				return c.JSON(
					http.StatusBadRequest,
					controllers.SetResponse(
						http.StatusBadRequest,
						"Parameters error",
						"Param ("+paramKey[k]+") is not a number",
					),
				)
			}
		}

		return next(c)
	}
}
