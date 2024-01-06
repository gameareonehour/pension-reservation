package release_note

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func GetLatestReleaseNotes(service *GetLatestReleaseNotesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response := service.GetLatestReleaseNotes()
		r, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		return ctx.JSON(r)
	}
}
