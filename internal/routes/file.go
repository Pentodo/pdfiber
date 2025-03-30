package fileRoutes

import (
	fileHandlers "PDFiber/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(api fiber.Router) {
	file := api.Group("/file")

	file.Post("/merge", fileHandlers.Merge)
	file.Post("/split", fileHandlers.Split)

}
