package main

import (
	"PDFiber/config"
	"PDFiber/router"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.Setup(app)

	os.MkdirAll(config.GlobalConfig.TempDir, os.ModePerm)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		sig := <-c
		log.Printf("Sinal recebido: %v\n", sig)

		os.RemoveAll(config.GlobalConfig.TempDir)
		os.Exit(0)
	}()

	addr := fmt.Sprintf(":%s", config.GlobalConfig.Port)
	log.Fatal(app.Listen(addr))
}
