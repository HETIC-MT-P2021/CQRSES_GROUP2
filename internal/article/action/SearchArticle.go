package action

import (
	"cqrses/internal/article/domain"
	"cqrses/internal/article/domain/entities"
	"cqrses/internal/article/responder"
	"cqrses/internal/pkg/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SearchArticleAction struct {
	Domain    *domain.ArticleService
	Responder service.Responder
}

func SearchArticle(c echo.Context) error {
	a := new(SearchArticleAction)
	a.Domain = domain.NewArticleService()
	a.Responder = responder.NewArticleResponder()

	return a.HandlerFunc(c)
}

func (a SearchArticleAction) HandlerFunc(c echo.Context) error {
	var articleData entities.ArticleData

	if err := c.Bind(&articleData); err != nil {
		payload := a.Domain.NewPayload(http.StatusBadRequest, err)
		return a.Responder.Response(c, payload)
	}

	objectID := c.Param("id")
	payload := a.Domain.Search(objectID)

	return a.Responder.Response(c, payload)
}
