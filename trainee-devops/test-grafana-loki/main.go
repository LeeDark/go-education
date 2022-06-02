package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/ping", pingHandler)
	e.Start(":8080")
}

func pingHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Answer string `json:"answer"`
	}{Answer: "pong"})
}
