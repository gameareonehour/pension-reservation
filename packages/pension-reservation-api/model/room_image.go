package model

import "gorm.io/gorm"

type RoomImage struct {
	gorm.Model
	// 部屋ID
	RoomID uint
	// 部屋紹介画像URL
	ImageURL string
}
