package main

import (
	"log"
	"net/http"
	"quickflow/config"
	"quickflow/internal/application/user"
	"quickflow/internal/infrastructure/database"
	"quickflow/internal/infrastructure/repository"
	"quickflow/internal/interfaces/httpserver"
	"quickflow/internal/interfaces/httpserver/handler"
	"quickflow/pkg/logger"

	"github.com/labstack/echo/v4"
)

func main() {
	err := logger.Init("info")
	if err != nil {
		panic(err)
	}

	// Initialize configuration
	logger.Info("Initialize configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	logger.Info("Initialize database connector...")
	database.InitDatabase(cfg)
	defer func() {
		sqlDB, _ := database.GetDB().DB()
		sqlDB.Close()
	}()

	// Initialize Echo instance
	logger.Info("Initialize HTTP IO...")
	e := echo.New()

	// Initialize User Application
	logger.Info("Initialize User application...")
	userRepo := repository.NewUserRepository(database.GetDB())
	userService := user.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Root handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "QuickFlow Main")
	})

	httpserver.SetupRoutes(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))

	logger.Info("Server started! Application launched.")
}
