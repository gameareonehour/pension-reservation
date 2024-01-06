package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// API responsible to deliver communication of
// application and client via http transport.
type API struct {
	core   *fiber.App
	router fiber.Router
}

func NewApi() *API {
	conf := fiber.Config{
		AppName: "pension-reservation",
	}

	api := fiber.New(conf)
	api.Use(recover.New())
	api.Use(logger.New())

	router := api.Group("/api").Group("/v1")
	return &API{core: api, router: router}
}

func (r *API) GetCore() *fiber.App {
	return r.core
}

func (r *API) GetRouter() fiber.Router {
	return r.router
}
