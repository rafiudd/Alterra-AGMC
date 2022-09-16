package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) HealthCheck(c echo.Context) error {
	result := h.controller.HealthCheck()
	return c.JSON(http.StatusOK, result)
}
