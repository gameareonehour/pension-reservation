package database

import (
	"pension-reservation-api/model"
	"time"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	err := createReleaseNotes(db)
	if err != nil {
		return err
	}

	return nil
}

func createReleaseNotes(db *gorm.DB) error {
	releaseNote := &model.ReleaseNote{
		Model: gorm.Model {
			CreatedAt: time.Date(2016, 2, 15, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2016, 2, 15, 0, 0, 0, 0, time.UTC),
		},
		Text: "サイトオープンしました。",
	}

	err := db.Create(releaseNote).Error
	if err != nil {
		return err
	}

	return nil
}
