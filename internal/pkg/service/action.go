package service

import (
	"github.com/labstack/echo/v4"
)

type Action interface {
	HandlerFunc(ctx echo.Context) error
}
