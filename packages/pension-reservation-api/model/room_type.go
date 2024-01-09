package model

import "gorm.io/gorm"

type RoomType struct {
	gorm.Model
	// 部屋タイプ名称
	Name string
}
