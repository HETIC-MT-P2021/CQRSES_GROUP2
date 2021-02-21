package service

import (
	"cqrses/pkg/payload"

	"github.com/labstack/echo/v4"
)

type Responder interface {
	Response(c echo.Context, payload payload.DomainPayload) error
}
