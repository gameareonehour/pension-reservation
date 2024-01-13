package catalog

import (
	"pension-reservation-api/core"
	"pension-reservation-api/openapi/generated"

	"github.com/gofiber/fiber/v2"
)

type Controller struct{
	svc *service
}

func NewController(svc *service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (c *Controller) GetCatalog(params generated.GetCatalogParams) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response, err := c.svc.GetCatalog(params)
		if err != nil {
			e, ok := err.(*core.IllegalArguments)
			if ok {
				return ctx.Status(fiber.StatusUnprocessableEntity).SendString(e.Error())
			}

			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}
