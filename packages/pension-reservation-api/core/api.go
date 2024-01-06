package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type API struct {
	instance *fiber.App
	router   fiber.Router
}

func NewAPI(l *Logger) *API {
	conf := fiber.Config{
		AppName:      "pension-reservation",
		ErrorHandler: onException(l),
	}

	api := fiber.New(conf)
	api.Use(logger.New())
	api.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	router := api.Group("/api").Group("/v1")
	return &API{instance: api, router: router}
}

func (r *API) Instance() *fiber.App {
	return r.instance
}

func (r *API) Router() fiber.Router {
	return r.router
}

func onException(l *Logger) func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		l.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("An error has occurred.")
	}
}
