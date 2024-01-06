package main

import (
	"context"
	"os"
	"os/signal"
	"pension-reservation-api/core"
	"pension-reservation-api/manipulation"
	"pension-reservation-api/mod/release_note"
	"syscall"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func main() {
	port := ":8080"

	db, err := core.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	a := core.NewApi()
	logger := core.NewLogger(os.Stdout)
	injector := do.New()

	provide(db, injector)

	go func() { _ = a.GetCore().Listen(port) }()

	var sig os.Signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig = <-c 

	logger.Printf("Signal received: %s",sig.String())
	logger.Printf("Shutting down app, waiting background process to finish")

	_ = a.GetCore().ShutdownWithContext(context.Background())
}

func provide(db *gorm.DB, injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (*release_note.GetLatestReleaseNotesService, error) {
		m := manipulation.NewGetLatestReleaseNotes(db)
		svc := release_note.NewGetLatestReleaseNotesService(m)

		return svc, nil
	})
}

func apiRoute(injector *do.Injector, logger *core.Logger) {
	// releaseNoteController := release_note.NewReleaseNoteController(injector, logger)

	// type si struct {}
}
