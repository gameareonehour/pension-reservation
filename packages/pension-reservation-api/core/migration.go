package core

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []interface{}{}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func DropTables(db *gorm.DB) error {
	return db.Migrator().DropTable()
}
