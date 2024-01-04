package release_note

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetLatestReleaseNotes(handler *GetLatestReleaseNotesService) echo.HandlerFunc {
	return func(c echo.Context) error {
		response := handler.GetLatestReleaseNotes()

		return c.JSON(http.StatusOK, response)
	}
}
