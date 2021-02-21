package service

import (
	"github.com/labstack/echo/v4"
)

type Responder interface {
	Response(c echo.Context, payload DomainPayload) error
}
