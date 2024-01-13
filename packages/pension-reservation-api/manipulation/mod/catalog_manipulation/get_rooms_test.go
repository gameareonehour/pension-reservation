package catalog_manipulation_test

import (
	"pension-reservation-api/core"
	"pension-reservation-api/manipulation/mod/catalog_manipulation"
	"pension-reservation-api/mod/catalog"
	"pension-reservation-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCatalog_GetRooms(t *testing.T) {
	setup := func() *gorm.DB {
		db, err := core.ConnectToInMemoryDatabase()
		if err != nil {
			t.Fatal(err)
		}

		err = core.Migrate(db)
		if err != nil {
			t.Fatal(err)
		}

		return db
	}

	cleanup := func(db *gorm.DB) {
		instance, err := db.DB()
		if err != nil {
			t.Fatal(err)
		}

		instance.Close()
	}

	t.Run("引数にnullを指定すると、部屋の全件取得できる", func(t *testing.T) {
		db := setup()

		for _, v := range data {
			err := db.Create(v).Error
			if err != nil {
				t.Fatal(err)
			}
		}

		defer cleanup(db)

		expected := &catalog.GetRoomsQueryResult{
			{
				ID:       1,
				Name:     "room1",
				Dayfee:   1,
				Type:     "type1",
				ImageURL: "/room1/image1",
			},
			{
				ID:       2,
				Name:     "room2",
				Dayfee:   2,
				Type:     "type2",
				ImageURL: "/room2/image1",
			},
		}

		query := catalog_manipulation.NewGetRooms(db)
		rooms, err := query.Run(nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, rooms)
	})

	t.Run("引数に部屋タイプを指定すると、指定した部屋タイプに一致する部屋を取得できる", func(t *testing.T) {
		db := setup()

		for _, v := range data {
			err := db.Create(v).Error
			if err != nil {
				t.Fatal(err)
			}
		}

		defer cleanup(db)

		expected := &catalog.GetRoomsQueryResult{
			{
				ID:       1,
				Name:     "room1",
				Dayfee:   1,
				Type:     "type1",
				ImageURL: "/room1/image1",
			},
		}

		query := catalog_manipulation.NewGetRooms(db)
		typeId := 1
		rooms, err := query.Run(&typeId)

		assert.Nil(t, err)
		assert.Equal(t, expected, rooms)
	})
}

var data = []interface{}{
	[]*model.RoomType{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "type1",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "type2",
		},
	},
	[]*model.Room{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name:        "room1",
			Description: "room1 description",
			Amenity:     "room1 amenity",
			Dayfee:      1,
			Capacity:    1,
			TypeID:      1,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name:        "room2",
			Description: "room2 description",
			Amenity:     "room2 amenity",
			Dayfee:      2,
			Capacity:    2,
			TypeID:      2,
		},
	},
	[]*model.RoomImage{
		{
			RoomID:   1,
			ImageURL: "/room1/image1",
		},
		{
			RoomID:   1,
			ImageURL: "/room1/image2",
		},
		{
			RoomID:   2,
			ImageURL: "/room2/image1",
		},
		{
			RoomID:   2,
			ImageURL: "/room2/image2",
		},
	},
}
