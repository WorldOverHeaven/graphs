package common

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
