package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func setEndpoints() *echo.Echo {
	e := echo.New()
	e.GET("/", home)
	e.GET("/ping", ping)
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
