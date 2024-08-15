package main

import (
	"net/http"
	"quickflow/internal/infrastructure/database"

	"github.com/labstack/echo/v4"
)

func main() {

	database.InitDatabase()
	defer func() {
		sqlDB, _ := database.GetDB().DB()
		sqlDB.Close()
	}()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
