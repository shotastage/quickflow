// File: pkg/errors/httperror.go

package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleHTTPError(c echo.Context, err error) error {
	var appErr *AppError
	if IsAppError(err) {
		// If it's an AppError, use a custom message and HTTP status code
		if As(err, &appErr) {
			return c.JSON(appErr.HTTPStatusCode(), map[string]string{
				"error":   appErr.Message,
				"details": err.Error(),
			})
		}
	} else {
		// If it's not an AppError, return 500 Internal Server Error
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "An unexpected error occurred",
		})
	}

	return c.NoContent(http.StatusInternalServerError)
}
