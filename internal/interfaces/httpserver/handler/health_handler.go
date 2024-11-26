// File: internal/interfaces/httpserver/handler/health_handler.go
package handler

import (
	"net/http"

	"quickflow/internal/application/health"
	"quickflow/pkg/errors"
	"quickflow/pkg/logger"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	healthService *health.HealthService
}

func NewHealthHandler(healthService *health.HealthService) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
}

// Handle processes the health check request
func (h *HealthHandler) Handle(c echo.Context) error {
	logger.Info("Health check requested")

	health, err := h.healthService.CheckHealth(c.Request().Context())
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return c.JSON(appErr.HTTPStatusCode(), health)
		}
		return c.JSON(http.StatusInternalServerError, health)
	}

	if health.Database.Status != "ok" {
		return c.JSON(http.StatusServiceUnavailable, health)
	}

	return c.JSON(http.StatusOK, health)
}
