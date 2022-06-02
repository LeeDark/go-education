package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setEndpoints() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", home)
	e.GET("/ping", ping)
	e.GET("/hello", hello)
	return e
}

// Active endpoint = HTTP Server gives HTML page
func home(c echo.Context) error {
	return c.HTML(
		http.StatusOK, `
		<HTML>
		<HEAD><TITLE>HTTP Server</TITLE></HEAD>
		<BODY><P align='center'>Hello! Welcome to HTTP Server written on Golang!</P></BODY>
		</HTML>`)
}

// Passive endpoint = HTTP Server gives JSON (XML) data = Frontend uses this JSON (XML) data
func ping(c echo.Context) error {
	//log.Println("Got ping")
	return c.JSON(http.StatusOK, struct {
		Answer string `json:"answer"`
	}{Answer: "pong"})
}

func hello(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		c.Error(echo.ErrBadRequest)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, struct {
		Answer string `json:"answer"`
	}{Answer: name})
}
