package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romankravchuk/toronto-wheels/internal/server/router/handlers/customers"
)

func NewV1() *fiber.App {
	r := fiber.New()
	r.Mount("/customers", newCustomersRouter())
	return r
}

func newCustomersRouter() *fiber.App {
	r := fiber.New()
	h := customers.New()
	r.Get("/", h.GetCustomers)
	r.Post("/", h.AddCustomer)
	r.Route("/:id", func(router fiber.Router) {
		router.Get("/", h.GetCustomer)
		router.Put("/", h.UpdateCustomer)
		router.Delete("/", h.DeleteCustomer)
	})
	return r
}
