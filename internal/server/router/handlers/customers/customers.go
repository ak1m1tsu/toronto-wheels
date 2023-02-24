package customers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romankravchuk/torolog"
	"github.com/romankravchuk/toronto-wheels/internal/server/router/handlers/common"
)

// CustomersHandler represents a handler for customers API.
type CustomersHandler struct {
	log *torolog.Logger
}

// New creates a new CustomersHandler instance with the provided torolog.
func New(log *torolog.Logger) *CustomersHandler {
	return &CustomersHandler{
		log: log,
	}
}

// TODO: make decision "How to implement GetCustomers?"

// GetCustomers returns the array of customers.
func (CustomersHandler) GetCustomers(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "get customers"})
}

// GetCustomer returns the customer details for the provided UUID.
// If the UUID is not found, it returns an error with a 404 status code.
func (h CustomersHandler) GetCustomer(ctx *fiber.Ctx) error {
	uuid, err := common.GetUUIDByFiberCtx(ctx)
	if err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "GetCustomer"},
		})
		return fiber.ErrNotFound
	}
	return ctx.JSON(fiber.Map{"message": "get customer " + uuid})
}

// UpdateCustomer updates the customer details for the provided UUID.
// If the UUID is not found, it returns an error with a 404 status code.
func (h CustomersHandler) UpdateCustomer(ctx *fiber.Ctx) error {
	uuid, err := common.GetUUIDByFiberCtx(ctx)
	if err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "UpdateCustomer"},
		})
		return fiber.ErrNotFound
	}
	return ctx.JSON(fiber.Map{"message": "update customer " + uuid})
}

// DeleteCustomer delets the customer details for the provided UUID.
// If the UUID is not found, it returns an error with a 404 status code.
func (h CustomersHandler) DeleteCustomer(ctx *fiber.Ctx) error {
	uuid, err := common.GetUUIDByFiberCtx(ctx)
	if err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "DeleteCustomer"},
		})
		return fiber.ErrNotFound
	}
	return ctx.JSON(fiber.Map{"message": "delete customer " + uuid})
}

// AddCustomer adds a new customer.
func (CustomersHandler) AddCustomer(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "add customer"})
}
