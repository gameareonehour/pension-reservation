package model

import "gorm.io/gorm"

type ReleaseNote struct {
	gorm.Model
	Text string `gorm:"column:text"`
}

func (ReleaseNote) TableName() string {
	return "release_notes"
}
