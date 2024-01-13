package core

import (
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToInMemoryDatabase() (*gorm.DB, error) {
	level := logger.Silent

	_, ok := os.LookupEnv("ENABLE_ORM_LOGGING")
	if ok {
		level = logger.Info
	} else {
		level = logger.Silent
	}

	return gorm.Open(sqlite.Open(":memory:?_pragma=foreign_keys(1)"), &gorm.Config{
		Logger: logger.Default.LogMode(level),
		NowFunc: func() time.Time {
			return time.Time{}
		},
	})
}

func RefreshInMemoryDatabase(db *gorm.DB) error {
	rows, err := db.Raw("select name from sqlite_master where type='table'").Rows()
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var t string

		err := rows.Scan(&t)
		if err != nil {
			return err
		}

		err = db.Exec("DELETE FROM " + t).Error
		if err != nil {
			return err
		}
	}

	return nil
}