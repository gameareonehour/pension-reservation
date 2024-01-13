package model

import "gorm.io/gorm"

type RoomImage struct {
	gorm.Model
	RoomID   uint   `gorm:"column:room_id"`   // 部屋ID
	ImageURL string `gorm:"column:image_url"` // 部屋紹介画像URL
}

func (RoomImage) TableName() string {
	return "room_images"
}
