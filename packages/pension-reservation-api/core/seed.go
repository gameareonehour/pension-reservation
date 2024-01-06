package core

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
	notes := []*model.ReleaseNote{
		{
			Model: gorm.Model{
				CreatedAt: time.Date(2016, 3, 3, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2016, 3, 3, 0, 0, 0, 0, time.UTC),
			},
			Text: "システムメンテナンスに伴うシステム障害のお詫びとご報告",
		},
		{
			Model: gorm.Model{
				CreatedAt: time.Date(2016, 3, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2016, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			Text: "３月２日にシステムメンテナンスを行います。",
		},
		{
			Model: gorm.Model{
				CreatedAt: time.Date(2016, 2, 29, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2016, 2, 29, 0, 0, 0, 0, time.UTC),
			},
			Text: "サイトデザインを更新しました。",
		},
		{
			Model: gorm.Model{
				CreatedAt: time.Date(2016, 2, 15, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2016, 2, 15, 0, 0, 0, 0, time.UTC),
			},
			Text: "サイトオープンしました。",
		},
	}

	err := db.Create(notes).Error
	if err != nil {
		return err
	}

	return nil
}
