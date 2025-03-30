package fileUtils

import (
	"PDFiber/config"
	"archive/zip"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePDFs(files []*multipart.FileHeader) (string, error) {
	savePath := filepath.Join(config.GlobalConfig.TempDir, uuid.New().String())
	err := os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("erro ao salvar arquivo temporário")
	}

	filePaths := make([]string, 0, len(files))
	for i, file := range files {
		fileName := fmt.Sprintf("%d.pdf", i)
		filePath := filepath.Join(savePath, fileName)

		err := saveFile(file, filePath)
		if err != nil {
			return "", fmt.Errorf("erro ao salvar arquivos temporários")
		}
		filePaths = append(filePaths, filePath)
	}

	mergedPDFPath := filepath.Join(savePath, "mergedFile.pdf")
	err = pdfcpu.MergeCreateFile(filePaths, mergedPDFPath, false, nil)
	if err != nil {
		return "", fmt.Errorf("erro ao juntar PDFs")
	}

	return mergedPDFPath, nil
}

func SplitAndZipPDF(file *multipart.FileHeader) (string, error) {
	savePath := filepath.Join(config.GlobalConfig.TempDir, uuid.New().String())
	err := os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("erro ao salvar arquivo temporário")
	}

	filePath := filepath.Join(savePath, "input.pdf")
	err = saveFile(file, filePath)
	if err != nil {
		return "", fmt.Errorf("erro ao salvar arquivo temporário")
	}

	splitPath := filepath.Join(savePath, "split")
	err = os.MkdirAll(splitPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("erro ao salvar arquivo temporário")
	}

	err = pdfcpu.SplitFile(filePath, splitPath, 1, nil)
	if err != nil {
		return "", fmt.Errorf("erro ao dividir o arquivo PDF")
	}

	zipFilePath := filepath.Join(savePath, "zippedFile.zip")
	err = createZipFromDirectory(splitPath, zipFilePath)
	if err != nil {
		return "", fmt.Errorf("erro ao criar arquivo zip")
	}

	return zipFilePath, nil
}

func saveFile(file *multipart.FileHeader, destination string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func createZipFromDirectory(sourceDir, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
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
}
