package release_note

import (
	"pension-reservation-api/core/api"
	"pension-reservation-api/core/logging"

	"github.com/samber/do"
)

type router struct {
	injector *do.Injector
	api      *api.API
	logger   *logging.Logger
}

func NewRouter(
	injector *do.Injector,
	api *api.API,
	logger *logging.Logger,
) *router {
	return &router{
		injector: injector,
		api:      api,
		logger:   logger,
	}
}

func (r *router) Serve() {
	getLatestReleaseNotesService := do.MustInvoke[*GetLatestReleaseNotesService](r.injector)

	{
		apiRouter := r.api.GetRouter().Group("/release-notes")
		apiRouter.Get("/", HandleGetLatestReleaseNotes(getLatestReleaseNotesService))
	}
}
