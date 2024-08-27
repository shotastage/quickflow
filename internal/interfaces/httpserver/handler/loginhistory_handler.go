// File: internal/interfaces/httpserver/handler/loginhistory_handler.go

package handler

import (
	"net/http"

	"quickflow/internal/application/loginhistory"
	"quickflow/pkg/errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LoginHistoryHandler struct {
	service loginhistory.Service
}

func NewLoginHistoryHandler(service loginhistory.Service) *LoginHistoryHandler {
	return &LoginHistoryHandler{service: service}
}

func (h *LoginHistoryHandler) RecordLogin(c echo.Context) error {
	userIDStr := c.Get("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	ipAddress := c.RealIP()
	userAgent := c.Request().UserAgent()

	if err := h.service.RecordLogin(c.Request().Context(), userID, ipAddress, userAgent); err != nil {
		return errors.HandleHTTPError(c, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *LoginHistoryHandler) GetUserLoginHistory(c echo.Context) error {
	userIDStr := c.Param("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	histories, err := h.service.GetUserLoginHistory(c.Request().Context(), userID)
	if err != nil {
		return errors.HandleHTTPError(c, err)
	}

	return c.JSON(http.StatusOK, histories)
}
