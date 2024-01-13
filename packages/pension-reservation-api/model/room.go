package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name        string `gorm:"column:name"`        // 部屋名称
	Description string `gorm:"column:description"` // 部屋概要
	Amenity     string `gorm:"column:amenity"`     // アメニティーテキスト
	Dayfee      int    `gorm:"column:dayfee"`      // 一泊料金
	Capacity    int    `gorm:"column:capacity"`    // 定員数

	// 部屋タイプ
	TypeID uint `gorm:"column:type_id"`
	Type   RoomType

	// 紹介画像
	Images []RoomImage
}

func (Room) TableName() string {
	return "rooms"
}
