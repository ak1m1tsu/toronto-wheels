package main

import (
	"log"
	"os"

	"github.com/romankravchuk/logger"
	"github.com/romankravchuk/toronto-wheels/internal/server"
)

// API entry point
func main() {
	logger := logger.NewWithLevel(os.Stdout, logger.DebugLevel)
	app := server.New(logger)
	log.Fatal(app.Listen(":8000"))
}
