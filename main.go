package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"quickflow/config"
	"quickflow/internal/application/user"
	"quickflow/internal/infrastructure/database"
	"quickflow/internal/infrastructure/repository"
	"quickflow/internal/interfaces/httpserver"
	"quickflow/internal/interfaces/httpserver/handler"
	"quickflow/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application failed to start: %v", err)
	}
}

func run() error {
	// Initialize logger
	if err := logger.Init("info"); err != nil {
		return err
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize database
	db, err := database.InitDatabase(cfg)
	if err != nil {
		return err
	}
	defer func() {
		if err := database.CloseDatabase(db); err != nil {
			logger.Error("Failed to close database connection", err)
		}
	}()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize application services
	userService := user.NewUserService(userRepo)

	// Initialize HTTP handlers
	userHandler := handler.NewUserHandler(userService)

	statusHandler := handler.NewStatusHandler()

	// Initialize Echo instance
	e := initializeEcho()

	// Setup routes
	httpserver.SetupRoutes(e, userHandler, statusHandler)

	// Start server
	return startServer(e, cfg.Server.Port)
}

func initializeEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

func startServer(e *echo.Echo, port int) error {
	// Start server in a goroutine
	go func() {
		if err := e.Start(":" + strconv.Itoa(port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}
