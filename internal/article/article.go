package article

import (
	"cqrses/internal/article/action"

	"github.com/labstack/echo/v4"
)

// SetArticleRoutes defines all article routes.
func SetArticleRoutes(g *echo.Group) {
	g.POST("/article", action.CreateArticle)
	g.POST("/article/:id", action.UpdateArticle)
	g.GET("/article/:id", action.SearchArticle)
}
