package main

import (
	"context"
	"os"
	"os/signal"
	"pension-reservation-api/core/api"
	"pension-reservation-api/core/database"
	"pension-reservation-api/core/logging"
	"pension-reservation-api/manipulation"
	"pension-reservation-api/mod/release_note"
	"syscall"

	"github.com/samber/do"
)

var logger = logging.NewLogger(os.Stdout)

func main() {
	port := ":8080"

	db, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	api := api.NewApi()
	injector := do.New()

	do.Provide(injector, func(i *do.Injector) (*release_note.GetLatestReleaseNotesService, error) {
		m := manipulation.NewGetLatestReleaseNotes(db)
		svc := release_note.NewGetLatestReleaseNotesService(m)

		return svc, nil
	})

	release_note.NewRouter(injector, api, logger).Serve()

	go func() { _ = api.GetCore().Listen(port) }()

	var sig os.Signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig = <-c 

	logger.Printf("Signal received: %s",sig.String())
	logger.Printf("Shutting down app, waiting background process to finish")

	_ = api.GetCore().ShutdownWithContext(context.Background())
}
