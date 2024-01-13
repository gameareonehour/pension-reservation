package model

import "gorm.io/gorm"

type RoomType struct {
	gorm.Model
	Name string `gorm:"column:name"` // 部屋タイプ名称
}

func (RoomType) TableName() string {
	return "room_types"
}
