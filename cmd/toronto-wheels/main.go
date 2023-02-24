package main

import (
	"log"
	"os"

	"github.com/romankravchuk/torolog"
	"github.com/romankravchuk/toronto-wheels/internal/server"
)

// API entry point
func main() {
	logger := torolog.NewWithLevel(os.Stdout, torolog.DebugLevel)
	app := server.New(logger)
	log.Fatal(app.Listen(":8000"))
}
