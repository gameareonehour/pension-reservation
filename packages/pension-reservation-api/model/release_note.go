package model

import "gorm.io/gorm"

type ReleaseNote struct {
	gorm.Model
	Text string
}
