package main

import (
	"context"
	"os"
	"os/signal"
	"pension-reservation-api/core"
	"pension-reservation-api/manipulation"
	"pension-reservation-api/mod/catalog"
	"pension-reservation-api/mod/release_note"
	"pension-reservation-api/openapi/generated"
	"pension-reservation-api/openapi/server"
	"syscall"

	"github.com/pkg/errors"
	"github.com/samber/do"
	"gorm.io/gorm"
)

const port = ":8080"

func main() {
	logger := core.NewLogger(os.Stdout)

	db, err := core.ConnectToDatabase()
	if err != nil {
		logger.Error(errors.WithStack(err))
		os.Exit(1)
	}

	injector := do.New()

	provide(db, injector, logger)

	api := core.NewAPI(logger)
	srv := server.New(injector)

	generated.RegisterHandlers(api.Router(), srv)

	go func() { _ = api.Instance().Listen(port) }()

	var sig os.Signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig = <-c

	logger.Printf("Signal received: %s", sig.String())
	logger.Printf("Shutting down app, waiting background process to finish")

	_ = api.Instance().ShutdownWithContext(context.Background())
	_ = injector.Shutdown()
}

func provide(db *gorm.DB, injector *do.Injector, logger *core.Logger) {
	// provide service instances.
	do.Provide(injector, func(i *do.Injector) (*release_note.GetLatestReleaseNotesService, error) {
		q := manipulation.NewReleaseNoteManipulation(db)
		svc := release_note.NewGetLatestReleaseNotesService(q)

		return svc, nil
	})

	// provide controller instances.
	do.Provide(injector, func(i *do.Injector) (*release_note.Controller, error) {
		controller := release_note.NewController()

		return controller, nil
	})

	do.Provide(injector, func(i *do.Injector) (*catalog.Controller, error) {
		query := manipulation.NewCatalogManipulation(db)
		svc := catalog.NewService(query)
		ctr := catalog.NewController(svc)

		return ctr, nil
	})
}
