package fileHandlers

import (
	"PDFiber/config"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func Merge(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erro ao processar os arquivos")
	}

	filePaths := make([]string, 0, len(form.File["pdfs"]))
	for _, file := range form.File["pdfs"] {
		fileName := fmt.Sprintf("%s.pdf", uuid.New().String())
		filePath := filepath.Join(config.GlobalConfig.TempDir, fileName)
		err := c.SaveFile(file, filePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Erro ao salvar arquivos temporários")
		}
		filePaths = append(filePaths, filePath)
	}

	mergedPDFName := fmt.Sprintf("%s.pdf", uuid.New().String())
	mergedPDFPath := filepath.Join(config.GlobalConfig.TempDir, mergedPDFName)
	err = api.MergeCreateFile(filePaths, mergedPDFPath, false, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erro ao juntar PDFs")
	}

	c.Set("Content-Disposition", "attachment; filename=file.pdf")
	return c.SendFile(mergedPDFPath)
}

func Split(c *fiber.Ctx) error {
	file, err := c.FormFile("pdf")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erro ao processar o arquivo PDF")
	}

	splitDir := filepath.Join(config.GlobalConfig.TempDir, uuid.New().String())
	os.MkdirAll(splitDir, os.ModePerm)

	fileName := fmt.Sprintf("%s.pdf", uuid.New().String())
	filePath := filepath.Join(config.GlobalConfig.TempDir, fileName)

	err = c.SaveFile(file, filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erro ao salvar arquivo temporário")
	}

	err = api.SplitFile(filePath, splitDir, 1, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erro ao dividir o arquivo PDF")
	}

	zipFilePath := filepath.Join(config.GlobalConfig.TempDir, "files.zip")
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erro ao criar arquivo zip")
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)

	err = filepath.Walk(splitDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		zipEntry, err := zipWriter.Create(info.Name())
		if err != nil {
			return err
		}
		fileContent, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fileContent.Close()

		_, err = io.Copy(zipEntry, fileContent)
		return err
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erro ao criar conteúdo do arquivo zip")
	}

	zipWriter.Close()

	c.Set("Content-Disposition", "attachment; filename=files.zip")
	return c.SendFile(zipFilePath)
}
