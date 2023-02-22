package router

import "github.com/gofiber/fiber/v2"

func NewV1() *fiber.App {
	r := fiber.New()
	return r
}
