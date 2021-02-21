package responder

import (
	"net/http"

	"cqrses/internal/pkg/response"
	"cqrses/pkg/payload"

	"github.com/labstack/echo/v4"
)

type ArticleResponder struct{}

func NewArticleResponder() *ArticleResponder {
	return new(ArticleResponder)
}

func (a ArticleResponder) Response(c echo.Context, payload payload.DomainPayload) error {
	var status int
	var result = payload.GetResult()
	var res *response.Response

	switch status := payload.GetStatus(); status {
	case http.StatusCreated:
		res = response.NewResponse(status, http.StatusText(status), result)
	default:
		status = http.StatusInternalServerError
		res = response.NewResponse(status, "Unknown DomainPayload Status", result)
	}

	return c.JSON(status, res)
}
