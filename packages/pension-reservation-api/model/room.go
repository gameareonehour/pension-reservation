package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	// 部屋名称
	Name string
	// 部屋概要
	Description string
	// アメニティーテキスト
	Amenity string
	// 一泊料金
	Dayfee int
	// 定員数
	Capacity int
	// 部屋タイプ
	TypeID *uint
	// 紹介画像
	Images []RoomImage
}
