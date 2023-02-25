package customers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/romankravchuk/torolog"
	"github.com/romankravchuk/toronto-wheels/internal/server/router/errors"
	"github.com/romankravchuk/toronto-wheels/internal/server/router/handlers/common"
	"github.com/romankravchuk/toronto-wheels/internal/server/router/handlers/models"
	"github.com/romankravchuk/toronto-wheels/internal/server/router/validator"
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
	return ctx.JSON(fiber.Map{"customers": []any{}})
}

// GetCustomer returns the customer details for the provided UUID.
// If the UUID is not found, it returns an error with a 404 status code.
func (h CustomersHandler) GetCustomer(ctx *fiber.Ctx) error {
	uuid, err := common.GetUUIDByFiberCtx(ctx)
	if err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Action", Value: "Get UUID from context"},
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "GetCustomer"},
		})
		return ctx.Status(http.StatusNotFound).SendString(errors.ErrGetUUIDFromContext.Error())
	}
	// atomic operation of getting the customer from the customer table
	return ctx.JSON(fiber.Map{"id": uuid})
}

// UpdateCustomer updates the customer details for the provided UUID.
// If the UUID is not found, it returns an error with a 404 status code.
func (h CustomersHandler) UpdateCustomer(ctx *fiber.Ctx) error {
	uuid, err := common.GetUUIDByFiberCtx(ctx)
	if err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Action", Value: "Get UUID from context"},
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "UpdateCustomer"},
		})
		return ctx.Status(http.StatusNotFound).SendString(errors.ErrGetUUIDFromContext.Error())
	}
	var payload models.CustomerPayload
	if err := ctx.BodyParser(&payload); err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Action", Value: "Parse body"},
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "UpdateCustomer"},
		})
		return ctx.Status(http.StatusBadRequest).SendString(errors.ErrBodyParse.Error())
	}
	if err := validator.Validate(payload); err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Action", Value: "Validate payload"},
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "UpdateCustomer"},
		})
		return ctx.Status(http.StatusBadRequest).SendString(errors.ErrValidationFailed.Error())
	}
	// check if such customer exists

	// atomic operation of updatting the customer in the customer table

	return ctx.JSON(fiber.Map{"id": uuid})
}

// DeleteCustomer delets the customer details for the provided UUID.
// If the UUID is not found, it returns an error with a 404 status code.
func (h CustomersHandler) DeleteCustomer(ctx *fiber.Ctx) error {
	uuid, err := common.GetUUIDByFiberCtx(ctx)
	if err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Action", Value: "Get UUID from context"},
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "UpdateCustomer"},
		})
		return ctx.Status(http.StatusNotFound).SendString(errors.ErrGetUUIDFromContext.Error())
	}
	return ctx.JSON(fiber.Map{"success": true, "id": uuid})
}

// AddCustomer adds a new customer.
func (h CustomersHandler) AddCustomer(ctx *fiber.Ctx) error {
	var payload models.CustomerPayload
	if err := ctx.BodyParser(&payload); err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Action", Value: "Parse body"},
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "UpdateCustomer"},
		})
		return ctx.Status(http.StatusBadRequest).SendString(errors.ErrBodyParse.Error())
	}
	if err := validator.Validate(payload); err != nil {
		h.log.Error("", err, torolog.Fields{
			{Key: "Action", Value: "Validate payload"},
			{Key: "Handler", Value: "CustomersHandler"},
			{Key: "HandlerFunc", Value: "UpdateCustomer"},
		})
		return ctx.Status(http.StatusBadRequest).SendString(errors.ErrValidationFailed.Error())
	}
	// check if such customer already exists

	// atomic operation of adding the customer to the customer table

	return ctx.JSON(fiber.Map{"customer": payload})
}
