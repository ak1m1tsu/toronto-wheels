package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romankravchuk/torolog"
	"github.com/romankravchuk/toronto-wheels/internal/server/router/handlers/customers"
)

func NewV1(log *torolog.Logger) *fiber.App {
	r := fiber.New()
	r.Mount("/customers", newCustomersRouter(log))
	return r
}

func newCustomersRouter(log *torolog.Logger) *fiber.App {
	r := fiber.New()
	h := customers.New(log)
	r.Get("/", h.GetCustomers)
	r.Post("/", h.AddCustomer)
	r.Route("/:id", func(router fiber.Router) {
		router.Get("/", h.GetCustomer)
		router.Put("/", h.UpdateCustomer)
		router.Delete("/", h.DeleteCustomer)
	})
	return r
}
