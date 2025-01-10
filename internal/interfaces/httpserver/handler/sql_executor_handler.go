// File: internal/interfaces/httpserver/handler/sql_executor_handler.go

package handler

import (
	"net/http"
	"quickflow/internal/application/sqlservice"
	"quickflow/internal/domain/sqlexecutor"
	"quickflow/pkg/errors"

	"github.com/labstack/echo/v4"
)

type SQLExecutorHandler struct {
	service sqlservice.SQLExecutorService
}

func NewSQLExecutorHandler(service sqlservice.SQLExecutorService) *SQLExecutorHandler {
	return &SQLExecutorHandler{
		service: service,
	}
}

func (h *SQLExecutorHandler) ExecuteQuery(c echo.Context) error {
	var req sqlexecutor.QueryRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	result, err := h.service.ExecuteQuery(c.Request().Context(), req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return c.JSON(appErr.HTTPStatusCode(), map[string]string{
				"error": appErr.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, result)
}
