package main

import (
	"context"
	"os"
	"os/signal"
	"pension-reservation-api/core"
	catalog_manipulation "pension-reservation-api/manipulation/catalog"
	release_note_manipulation "pension-reservation-api/manipulation/release_note"
	"pension-reservation-api/mod/catalog"
	"pension-reservation-api/mod/release_note"
	"pension-reservation-api/openapi/generated"
	"pension-reservation-api/openapi/server"
	"syscall"

	"github.com/pkg/errors"
	"go.uber.org/dig"
	// "github.com/samber/do"
)

const port = ":8080"

func main() {
	logger := core.NewLogger(os.Stdout)

	db, err := core.ConnectToDatabase()
	if err != nil {
		logger.Error(errors.WithStack(err))
		os.Exit(1)
	}

	container := dig.New()

	// provide release note controller.
	container.Provide(func () *release_note.Controller {
		query := release_note_manipulation.NewQuery(db)
		svc := release_note.NewService(query)
		ctr := release_note.NewController(svc)

		return ctr
	})

	// provide catalog controller.
	container.Provide(func () *catalog.Controller {
		query := catalog_manipulation.NewQuery(db)
		svc := catalog.NewService(query)
		ctr := catalog.NewController(svc)

		return ctr
	})

	api := core.NewAPI(logger)
	srv := server.New(container)

	generated.RegisterHandlers(api.Router(), srv)

	go func() { _ = api.Instance().Listen(port) }()

	var sig os.Signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig = <-c

	logger.Printf("Signal received: %s", sig.String())
	logger.Printf("Shutting down app, waiting background process to finish")

	_ = api.Instance().ShutdownWithContext(context.Background())
}
