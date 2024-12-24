// File: internal/interfaces/httpserver/handler/table_handler.go
package handler

import (
	"net/http"
	"quickflow/internal/application/table"
	"quickflow/internal/domain/tableentity"
	"quickflow/pkg/errors"

	"github.com/labstack/echo/v4"
)

type TableHandler struct {
	service *table.TableService
}

func NewTableHandler(service *table.TableService) *TableHandler {
	return &TableHandler{service: service}
}

func (h *TableHandler) CreateTable(c echo.Context) error {
	var table tableentity.Table
	if err := c.Bind(&table); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if err := h.service.CreateTable(c.Request().Context(), &table); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return c.JSON(appErr.HTTPStatusCode(), map[string]string{
				"error": appErr.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Table created successfully",
		"table":   table,
	})
}
