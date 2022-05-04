package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setEndpoints() *gin.Engine {
	// FIXME: use Release mode for benchmarking
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	// FIXME: don't use Logger middleware for benchmarking
	e.Use(gin.Recovery())
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
	//log.Println("Got ping")
	// c.JSON(http.StatusOK, struct {
	// 	Answer string `json:"answer"`
	// }{Answer: "pong"})
	c.JSON(http.StatusOK, gin.H{"answer": "pong"})
}
