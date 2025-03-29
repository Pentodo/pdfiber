package main

import (
	_ "PDFiber/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title PDFiber
// @BasePath /api
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/api",
		func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

	app.Listen(":3000")
}
