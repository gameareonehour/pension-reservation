package server

import (
	"pension-reservation-api/mod/catalog"
	"pension-reservation-api/mod/release_note"
	"pension-reservation-api/openapi/generated"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type server struct {
	container *dig.Container
}

var _ generated.ServerInterface = (*server)(nil)

func New(container *dig.Container) *server {
	return &server{
		container: container,
	}
}

func (s *server) GetReleaseNotes(ctx *fiber.Ctx) error {
	var ctr *release_note.Controller

	if err := s.container.Invoke(func(o *release_note.Controller) {
		ctr = o
	}); err != nil {
		return err
	}

	return ctr.GetLatestReleaseNotes()(ctx)
}

func (s *server) GetCatalog(ctx *fiber.Ctx, params generated.GetCatalogParams) error {
	var ctr *catalog.Controller

	if err := s.container.Invoke(func(o *catalog.Controller) {
		ctr = o
	}); err != nil {
		return err
	}

	return ctr.GetCatalog(params)(ctx)
}

func (s *server) PostVacancyRoomsSearch(ctx *fiber.Ctx) error {
	return nil
}
