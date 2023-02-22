package main

import (
	"log"

	"github.com/romankravchuk/toronto-wheels/internal/server"
)

// API entry point
func main() {
	app := server.New()
	log.Fatal(app.Listen(":8000"))
}
