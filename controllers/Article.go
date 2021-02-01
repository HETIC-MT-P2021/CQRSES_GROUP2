package controllers

import (
	"cqrses/domain"
	"cqrses/domain/article"
	"cqrses/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

// CreateArticle create an article
func CreateArticle(c echo.Context) error {
	var articleStore models.ArticleStore

	if err := c.Bind(&articleStore); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Validation error", err.Error()))
	}

	if err := validation.ValidateStruct(&articleStore,
		validation.Field(&articleStore.Title, validation.Required),
		validation.Field(&articleStore.Content, validation.Required),
	); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Validation error", err.Error()))
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID := claims["id"].(float64)

	articleStore.AuthorID = uint(authorID)

	command := article.CreateArticleCommand{
		ArticleStore: articleStore,
	}
	res, err := domain.Cb.Dispatch(command)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Article error", err.Error()))
	}

	return c.JSON(http.StatusCreated, SetResponse(http.StatusCreated, "Article created", res))
}
