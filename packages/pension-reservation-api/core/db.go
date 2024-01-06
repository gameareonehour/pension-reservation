package core

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase() (*gorm.DB, error) {
	username, err := Env("DB_USERNAME")
	if err != nil {
		return nil, err
	}

	password, err := Env("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	host, err := Env("DB_HOST")
	if err != nil {
		return nil, err
	}

	port, err := Env("DB_PORT")
	if err != nil {
		return nil, err
	}

	dbname, err := Env("DB_NAME")
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}