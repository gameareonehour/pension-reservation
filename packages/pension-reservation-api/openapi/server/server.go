package server

import (
	"pension-reservation-api/mod/catalog"
	"pension-reservation-api/mod/release_note"
	"pension-reservation-api/openapi/generated"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type server struct {
	injector *do.Injector
}

var _ generated.ServerInterface = (*server)(nil)

func New(injector *do.Injector) *server {
	return &server{
		injector: injector,
	}
}

func (s *server) GetReleaseNotes(ctx *fiber.Ctx) error {
	controller := do.MustInvoke[*release_note.Controller](s.injector)
	service := do.MustInvoke[*release_note.GetLatestReleaseNotesService](s.injector)

	return controller.GetLatestReleaseNotes(service)(ctx)
}

func (s *server) GetCatalog(ctx *fiber.Ctx, params generated.GetCatalogParams) error {
	ctr := do.MustInvoke[*catalog.Controller](s.injector)

	return ctr.GetCatalog(params)(ctx)
}

func (s *server) PostVacancyRoomsSearch(ctx *fiber.Ctx) error {
	return nil
}
