package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func setEndpoints() *fiber.App {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Get("/", home)
	app.Get("/ping", ping)
	return app
}

// Active endpoint = HTTP Server gives HTML page
func home(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).Render("index", nil)
}

// Passive endpoint = HTTP Server gives JSON (XML) data = Frontend uses this JSON (XML) data
func ping(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"answer": "pong"})
}
