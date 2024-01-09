package core

import (
	"pension-reservation-api/model"
	"time"

	"gorm.io/gorm"
)

const (
	roomTypeJapanese = 1
	roomTypeWestern = 2
	roomTypeJapaneseWestern = 3
)

func Seed(db *gorm.DB) error {
	s := seeder{
		db: db,
	}

	err := s.createReleaseNotes()
	if err != nil {
		return err
	}

	err = s.createRoomTypes()
	if err != nil {
		return err
	}
	
	err = s.createRooms()
	if err != nil {
		return err
	}

	err = s.associateRoomTypes()
	if err != nil {
		return err
	}

	return nil
}

type seeder struct {
	db *gorm.DB
}

func (s *seeder) createReleaseNotes() error {
	notes := []*model.ReleaseNote{
		{
			Model: gorm.Model{
				CreatedAt: time.Date(2016, 3, 3, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2016, 3, 3, 0, 0, 0, 0, time.UTC),
			},
			Text: "システムメンテナンスに伴うシステム障害のお詫びとご報告",
		},
		{
			Model: gorm.Model{
				CreatedAt: time.Date(2016, 3, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2016, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			Text: "３月２日にシステムメンテナンスを行います。",
		},
		{
			Model: gorm.Model{
				CreatedAt: time.Date(2016, 2, 29, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2016, 2, 29, 0, 0, 0, 0, time.UTC),
			},
			Text: "サイトデザインを更新しました。",
		},
		{
			Model: gorm.Model{
				CreatedAt: time.Date(2016, 2, 15, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2016, 2, 15, 0, 0, 0, 0, time.UTC),
			},
			Text: "サイトオープンしました。",
		},
	}

	err := s.db.Create(notes).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *seeder) createRoomTypes() error {
	types := []model.RoomType {
		{
			Model: gorm.Model{
				ID: roomTypeJapanese,
			},
			Name: "和室",
		},
		{
			Model: gorm.Model{
				ID: roomTypeWestern,
			},
			Name: "洋室",
		},
		{
			Model: gorm.Model{
				ID: roomTypeJapaneseWestern,
			},
			Name: "和洋室",
		},
	}

	err := s.db.Create(types).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *seeder) createRooms() error {
	rooms := []model.Room {
		{
			Name: "ゆとりの和室",
			Description: "お風呂・トイレも部屋内にある、広めの和室です。部活・サークルなど、気の合う仲間たちと大人数で利用するのに適しています。",
			Amenity: "テレビ、エアコン、冷蔵庫、部屋着、ドライヤー、シャンプー、リンス",
			Dayfee: 9000,
			Capacity: 4,
			TypeID: UIntPtr(roomTypeJapanese),
		},
		{
			Name: "ゆとりの和室",
			Description: "お風呂・トイレも部屋内にある、広めの和室です。部活・サークルなど、気の合う仲間たちと大人数で利用するのに適しています。",
			Amenity: "テレビ、エアコン、冷蔵庫、部屋着、ドライヤー、シャンプー、リンス",
			Dayfee: 12000,
			Capacity: 6,
			TypeID: UIntPtr(roomTypeJapanese),
		},
		{
			Name: "ゆとりの和室",
			Description: "お風呂・トイレも部屋内にある、広めの和室です。部活・サークルなど、気の合う仲間たちと大人数で利用するのに適しています。",
			Amenity: "テレビ、エアコン、冷蔵庫、部屋着、ドライヤー、シャンプー、リンス",
			Dayfee: 21000,
			Capacity: 8,
			TypeID: UIntPtr(roomTypeJapanese),
		},
		{
			Name: "落ち着きのある洋室",
			Description: "２〜３名でのご利用に最適な落ち着きのある洋室です。ファミリー様も大歓迎です。トイレは部屋内にありますが、お風呂は大浴場利用となります。",
			Amenity: "テレビ、エアコン、冷蔵庫、部屋着",
			Dayfee: 8000,
			Capacity: 3,
			TypeID: UIntPtr(roomTypeWestern),
		},
		{
			Name: "落ち着きのある洋室",
			Description: "２〜３名でのご利用に最適な落ち着きのある洋室です。ファミリー様も大歓迎です。トイレは部屋内にありますが、お風呂は大浴場利用となります。",
			Amenity: "テレビ、エアコン、冷蔵庫、部屋着",
			Dayfee: 10000,
			Capacity: 4,
			TypeID: UIntPtr(roomTypeWestern),
		},
		{
			Name: "落ち着きのある洋室",
			Description: "２〜３名でのご利用に最適な落ち着きのある洋室です。ファミリー様も大歓迎です。トイレは部屋内にありますが、お風呂は大浴場利用となります。",
			Amenity: "テレビ、エアコン、冷蔵庫、部屋着",
			Dayfee: 15000,
			Capacity: 6,
			TypeID: UIntPtr(roomTypeWestern),
		},
		{
			Name: "みんなで和洋室",
			Description: "最大６名まで利用可能な和洋室です。ベッド3台、布団3組の形となります。シャワー、トイレは部屋内にありますが、お風呂は大浴場利用となります。",
			Amenity: "テレビ、エアコン、冷蔵庫、部屋着",
			Dayfee: 10000,
			Capacity: 6,
			TypeID: UIntPtr(roomTypeJapaneseWestern),
		},
		{
			Name: "みんなで和洋室",
			Description: "最大６名まで利用可能な和洋室です。ベッド3台、布団3組の形となります。シャワー、トイレは部屋内にありますが、お風呂は大浴場利用となります。",
			Amenity: "テレビ、エアコン、冷蔵庫、部屋着",
			Dayfee: 15000,
			Capacity: 8,
			TypeID: UIntPtr(roomTypeJapaneseWestern),
		},
		{
			Name: "みんなで和洋室",
			Description: "最大６名まで利用可能な和洋室です。ベッド3台、布団3組の形となります。シャワー、トイレは部屋内にありますが、お風呂は大浴場利用となります。",
			Amenity: "テレビ、エアコン、冷蔵庫、部屋着",
			Dayfee: 32000,
			Capacity: 12,
			TypeID: UIntPtr(roomTypeJapaneseWestern),
		},
	}

	err := s.db.Create(rooms).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *seeder) associateRoomTypes() error {
	japaneseRooms := []*model.Room{}
	westernRooms := []*model.Room{}
	japaneseWesternRooms := []*model.Room{}

	err := s.db.Where(&model.Room{TypeID: UIntPtr(roomTypeJapanese)}).Find(&japaneseRooms).Error
	if err != nil {
		return err
	}

	err = s.db.Where(&model.Room{TypeID: UIntPtr(roomTypeWestern)}).Find(&westernRooms).Error
	if err != nil {
		return err
	}

	err = s.db.Where(&model.Room{TypeID: UIntPtr(roomTypeJapaneseWestern)}).Find(&japaneseWesternRooms).Error
	if err != nil {
		return err
	}

	for _, r := range japaneseRooms {
		roomImages := []model.RoomImage{
			{
				RoomID: r.ID,
				ImageURL: "",
			},
		}

		err := s.db.Create(roomImages).Error
		if err != nil {
			return err
		}
	}

	return nil
}
