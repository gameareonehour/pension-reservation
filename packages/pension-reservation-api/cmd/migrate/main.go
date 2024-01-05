package main

import (
	"os"
	"pension-reservation-api/core"
	"pension-reservation-api/pkg/logging"
)

func main() {
	logger := logging.NewLogger(os.Stdout)

	db, err := core.ConnectToDatabase()
	if err != nil {
		logger.Printf("unexpected error occurred during connect database: %+v", err)
		os.Exit(1)
	}

	err = core.Migrate(db)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}