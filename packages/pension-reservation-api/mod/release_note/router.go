package release_note

import (
	"pension-reservation-api/core/api"
	"pension-reservation-api/core/logging"
)

type router struct {
	api    *api.API
	logger *logging.Logger
	svc    *services
}

type services struct {
	getLatestReleaseNotes *GetLatestReleaseNotesService
}

func NewRouter(
	api    *api.API,
	logger *logging.Logger,
	getLatestReleaseNotes *GetLatestReleaseNotesService,
) *router {
	return &router {
		api: api,
		logger: logger,
		svc: &services{
			getLatestReleaseNotes: getLatestReleaseNotes,
		},
	}
}

func (r *router) Serve() {
	{
		apiRouter := r.api.GetRouter().Group("/release-notes")
		apiRouter.Get("/", GetLatestReleaseNotes(r.svc.getLatestReleaseNotes))
	}
}
