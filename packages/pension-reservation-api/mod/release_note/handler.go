package release_note

import (
	"github.com/gofiber/fiber/v2"
)

func HandleGetLatestReleaseNotes(service *GetLatestReleaseNotesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response := service.GetLatestReleaseNotes()

		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}
