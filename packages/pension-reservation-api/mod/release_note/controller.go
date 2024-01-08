package release_note

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetLatestReleaseNotes(service *GetLatestReleaseNotesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response, err := service.GetLatestReleaseNotes()
		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}
