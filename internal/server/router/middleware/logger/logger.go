package logger

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/romankravchuk/logger"
)

type Config struct {
	Next func(ctx *fiber.Ctx) bool

	Logger *logger.Logger
}

func New(cfg Config) fiber.Handler {
	var sublogger *logger.Logger
	if cfg.Logger == nil {
		sublogger = logger.New(os.Stdout)
	} else {
		sublogger = cfg.Logger
	}
	return func(ctx *fiber.Ctx) error {
		start := time.Now()
		if cfg.Next != nil && cfg.Next(ctx) {
			return ctx.Next()
		}
		err := ctx.Next()
		if err != nil {
			sublogger.Error("[TORONTO WHEELS]", err, logger.Fields{
				{Key: "Path", Value: ctx.Path()},
				{Key: "Mathod", Value: ctx.Method()},
				{Key: "Duration", Value: time.Since(start)},
			})
		} else {
			sublogger.Info("[TORONTO WHEELS]", logger.Fields{
				{Key: "Path", Value: ctx.Path()},
				{Key: "Method", Value: ctx.Method()},
				{Key: "Duration", Value: time.Since(start)},
			})
		}
		return err
	}
}
