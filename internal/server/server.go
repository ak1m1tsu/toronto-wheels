package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/romankravchuk/toronto-wheels/internal/server/router"
)

type Server struct {
	*fiber.App
}

func New() *Server {
	s := &Server{App: fiber.New()}
	s.MountDefaultMiddlewares()
	s.Route("/api", func(r fiber.Router) {
		r.Mount("/v1", router.NewV1())
	})
	return s
}

func (s *Server) MountDefaultMiddlewares() {
	s.Use(recover.New())
	s.Use(logger.New())
	s.Use(cors.New())
}
