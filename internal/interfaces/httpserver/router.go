// File: internal/interfaces/http/routes.go

package httpserver

import (
	"quickflow/internal/interfaces/httpserver/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	// User routes
	userGroup := e.Group("/users")
	{
		userGroup.POST("", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.PUT("/:id/password", userHandler.UpdatePassword)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}
}
