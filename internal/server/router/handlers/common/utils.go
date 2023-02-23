package common

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var errUUIDParam = errors.New("wrong value UUID")

func GetUUIDByFiberCtx(ctx *fiber.Ctx) (string, error) {
	param := ctx.Params("id")
	if !isValidUUID(param) {
		return "", errUUIDParam
	}
	return param, nil
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
