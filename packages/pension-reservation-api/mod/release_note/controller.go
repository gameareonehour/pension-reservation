package release_note

import (
	"pension-reservation-api/core"
	"pension-reservation-api/openapi/generated"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type controller struct {
	injector *do.Injector
	logger   *core.Logger
}

func NewReleaseNoteController(injector *do.Injector, logger *core.Logger) *controller {
	return &controller{
		injector: injector,
		logger: logger,
	}
}

func (c *controller) GetLatestReleaseNotes() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		service := do.MustInvoke[*GetLatestReleaseNotesService](c.injector)
		response := generated.GetReleaseNotesResponse{}

		notes := service.GetLatestReleaseNotes()

		for _, v := range notes {
			response.Items = append(response.Items, GetReleaseNotesResponseItem{
				Id: v.ID,
				Text: v.Text,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}
