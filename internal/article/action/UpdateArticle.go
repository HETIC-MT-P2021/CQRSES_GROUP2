package action

import (
	"cqrses/internal/article/domain"
	"cqrses/internal/article/domain/entities"
	"cqrses/internal/article/responder"
	"cqrses/internal/pkg/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UpdateArticleAction struct {
	Domain    *domain.ArticleService
	Responder service.Responder
}

// UpdateArticle update an article
func UpdateArticle(c echo.Context) error {
	a := new(UpdateArticleAction)
	a.Domain = domain.NewArticleService()
	a.Responder = responder.NewArticleResponder()

	return a.HandlerFunc(c)
}

func (u UpdateArticleAction) HandlerFunc(c echo.Context) error {
	var articleData entities.ArticleData

	objectID := c.Param("id")

	if err := c.Bind(&articleData); err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID := claims["id"].(float64)

	articleData.CreatedBy = uint(authorID)

	payload := u.Domain.Update(articleData, objectID)

	return u.Responder.Response(c, payload)
}
