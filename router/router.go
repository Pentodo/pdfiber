package router

import (
	fileRoutes "PDFiber/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	api := app.Group("/api", logger.New())

	fileRoutes.Setup(api)
}
