package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setEndpoints() *gin.Engine {
	// FIXME: use Release mode for benchmarking
	//gin.SetMode(gin.ReleaseMode)

	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery())

	e.LoadHTMLGlob("templates/*")

	e.GET("/", home)
	e.GET("/ping", ping)
	return e
}

// Active endpoint = HTTP Server gives HTML page
func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

// Passive endpoint = HTTP Server gives JSON (XML) data = Frontend uses this JSON (XML) data
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"answer": "pong"})
}
