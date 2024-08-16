package handler

import (
	"net/http"

	"quickflow/pkg/logger"

	"github.com/labstack/echo/v4"
)

// HealthHandler handles health check requests
type HealthHandler struct{}

// NewHealthHandler creates a new HealthHandler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Handle processes the health check request
func (h *HealthHandler) Handle(c echo.Context) error {
	logger.Info("Health check requested")

	response := map[string]string{
		"status":  "ok",
		"message": "Server is running",
	}

	return c.JSON(http.StatusOK, response)
}
