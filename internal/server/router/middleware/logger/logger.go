package logger

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/romankravchuk/torolog"
)

type Config struct {
	Next func(ctx *fiber.Ctx) bool

	Logger *torolog.Logger
}

func New(cfg Config) fiber.Handler {
	var sublogger *torolog.Logger
	if cfg.Logger == nil {
		sublogger = torolog.New(os.Stdout)
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
			sublogger.Error("[TORONTO WHEELS]", err, torolog.Fields{
				{Key: "Path", Value: ctx.Path()},
				{Key: "Mathod", Value: ctx.Method()},
				{Key: "Params", Value: ctx.AllParams()},
				{Key: "Duration", Value: time.Since(start)},
			})
		} else {
			sublogger.Info("[TORONTO WHEELS]", torolog.Fields{
				{Key: "Path", Value: ctx.Path()},
				{Key: "Method", Value: ctx.Method()},
				{Key: "Params", Value: ctx.AllParams()},
				{Key: "Duration", Value: time.Since(start)},
			})
		}
		return err
	}
}
