package controllers

import (
	"cqrses/cqrs"
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
	var articleData models.ArticleData

	if err := c.Bind(&articleData); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Validation error", err.Error()))
	}

	if err := validation.ValidateStruct(&articleData,
		validation.Field(&articleData.Title, validation.Required),
		validation.Field(&articleData.Content, validation.Required),
	); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Validation error", err.Error()))
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID := claims["id"].(float64)

	articleData.CreatedBy = uint(authorID)

	command := article.CreateArticleCommand{
		ArticleData: articleData,
	}

	cmdDescriptor := cqrs.NewCommandMessage(&command)
	res, err := domain.Cb.Dispatch(cmdDescriptor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Article error", err.Error()))
	}

	return c.JSON(http.StatusCreated, SetResponse(http.StatusCreated, "Article created", res))
}

// UpdateArticle update an article
func UpdateArticle(c echo.Context) error {
	aggregateID := c.Param("id")
	var articleData models.ArticleData

	if err := c.Bind(&articleData); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Validation error", err.Error()))
	}

	if err := validation.ValidateStruct(&articleData,
		validation.Field(&articleData.Title, validation.Required),
		validation.Field(&articleData.Content, validation.Required),
	); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Validation error", err.Error()))
	}

	command := article.EditArticleCommand{
		AggregateID: aggregateID,
		ArticleData: articleData,
	}
	cmdDescriptor := cqrs.NewCommandMessage(&command)
	res, err := domain.Cb.Dispatch(cmdDescriptor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Article error", err.Error()))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "Article updated", res))
}
