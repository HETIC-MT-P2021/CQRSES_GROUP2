package action

import (
	"cqrses/internal/article/domain"
	"cqrses/internal/article/domain/entities"
	"cqrses/internal/article/responder"
	"cqrses/internal/pkg/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type CreateArticleAction struct {
	Domain    *domain.ArticleService
	Responder service.Responder
}

func CreateArticle(c echo.Context) error {
	a := new(CreateArticleAction)
	a.Domain = domain.NewArticleService()
	a.Responder = responder.NewArticleResponder()

	return a.HandlerFunc(c)
}

func (a CreateArticleAction) HandlerFunc(c echo.Context) error {
	var articleData entities.ArticleData

	if err := c.Bind(&articleData); err != nil {
		payload := a.Domain.NewPayload(http.StatusBadRequest, err)
		return a.Responder.Response(c, payload)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID := claims["id"].(float64)

	articleData.CreatedBy = uint(authorID)

	payload := a.Domain.Create(articleData)

	return a.Responder.Response(c, payload)
}
