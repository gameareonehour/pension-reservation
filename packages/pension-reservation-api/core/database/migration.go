package database

import (
	"pension-reservation-api/core/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []interface{}{
		&model.ReleaseNote{},
	}

	for _, m := range models {
		err := db.AutoMigrate(m)
		if err != nil {
			return err
		}
	}

	return nil
}

func DropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&model.ReleaseNote{},
	)
}
