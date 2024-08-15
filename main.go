package main

import (
	"net/http"
	"quickflow/internal/application/user"
	"quickflow/internal/infrastructure/database"
	"quickflow/internal/infrastructure/repository"
	"quickflow/internal/interfaces/httpserver"
	"quickflow/internal/interfaces/httpserver/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDatabase()
	defer func() {
		sqlDB, _ := database.GetDB().DB()
		sqlDB.Close()
	}()

	e := echo.New()

	// Initialize User Application
	userRepo := repository.NewUserRepository(database.GetDB())
	userService := user.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	httpserver.SetupRoutes(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
