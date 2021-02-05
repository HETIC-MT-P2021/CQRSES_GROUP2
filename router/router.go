package router

import (
	"cqrses/controllers"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// InitRoutes initiates all routes.
func InitRoutes(e *echo.Echo) {
	// Public routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, controllers.SetResponse(http.StatusOK, "Welcome in "+viper.GetString("APP_NAME"), e.Routes()))
	})

	// Authentication routes
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)

	// Restricted group
	r := e.Group("/api", middleware.JWT([]byte("secret")))

	r.POST("/article", controllers.CreateArticle)
	r.POST("/article/:id", controllers.UpdateArticle)

	// Refresh token routes
	r.POST("/refresh", controllers.RefreshToken)

}
