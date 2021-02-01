package controllers

import (
	"cqrses/domain"
	"cqrses/domain/article"
	"cqrses/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateArticle create an article
func CreateArticle(c echo.Context) error {
	var articleStore models.ArticleStore

	command := article.CreateArticleCommand{
		ArticleStore: articleStore,
	}
	res, err := domain.Cb.Dispatch(command)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Article error", err.Error()))
	}

	return c.JSON(http.StatusCreated, SetResponse(http.StatusCreated, "Article created", res))
}
