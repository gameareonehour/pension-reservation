package main

import (
	"os"
	"pension-reservation-api/core/database"
	"pension-reservation-api/core/logging"
)

func main() {
	logger := logging.NewLogger(os.Stdout)

	db, err := database.ConnectToDatabase()
	if err != nil {
		logger.Printf("unexpected error occurred during connect database: %+v", err)
		os.Exit(1)
	}

	err = database.Seed(db)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
