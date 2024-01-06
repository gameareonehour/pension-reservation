package release_note

import (
	"pension-reservation-api/openapi/generated"

	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetLatestReleaseNotes(service *GetLatestReleaseNotesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response := generated.GetReleaseNotesResponse{}

		notes, err := service.GetLatestReleaseNotes()
		if err != nil {
			return err
		}

		for _, v := range notes {
			response.Items = append(response.Items, GetReleaseNotesResponseItem{
				Id:        v.ID,
				Text:      v.Text,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}
