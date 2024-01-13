package release_note

import (
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

func (c *Controller) GetLatestReleaseNotes() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response, err := c.svc.GetLatestReleaseNotes()
		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}
