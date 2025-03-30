package fileHandlers

import (
	fileUtils "PDFiber/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func Merge(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erro ao processar os arquivos")
	}

	mergedPDFPath, err := fileUtils.MergePDFs(form.File["pdfs"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	c.Set("Content-Disposition", "attachment; filename=file.pdf")
	return c.SendFile(mergedPDFPath)
}

func Split(c *fiber.Ctx) error {
	file, err := c.FormFile("pdf")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erro ao processar o arquivo PDF")
	}

	zipFilePath, err := fileUtils.SplitAndZipPDF(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	c.Set("Content-Disposition", "attachment; filename=files.zip")
	return c.SendFile(zipFilePath)
}
