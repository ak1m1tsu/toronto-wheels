package customers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romankravchuk/toronto-wheels/internal/server/router/handlers/common"
)

type CustomersHandler struct{}

func New() *CustomersHandler {
	return &CustomersHandler{}
}

// TODO: make decision "How to implement GetCustomers?"
func (CustomersHandler) GetCustomers(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "get customers"})
}

func (h CustomersHandler) GetCustomer(ctx *fiber.Ctx) error {
	uuid, err := common.GetUUIDByFiberCtx(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{"message": "get customer " + uuid})
}

func (CustomersHandler) UpdateCustomer(ctx *fiber.Ctx) error {
	uuid, err := common.GetUUIDByFiberCtx(ctx)
	if err != nil {
		return ctx.SendString(fiber.ErrNotFound.Error())
	}
	return ctx.JSON(fiber.Map{"message": "update customer " + uuid})
}

func (CustomersHandler) DeleteCustomer(ctx *fiber.Ctx) error {
	uuid, err := common.GetUUIDByFiberCtx(ctx)
	if err != nil {
		return ctx.SendString(fiber.ErrNotFound.Error())
	}
	return ctx.JSON(fiber.Map{"message": "delete customer " + uuid})
}

func (CustomersHandler) AddCustomer(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "add customer"})
}
