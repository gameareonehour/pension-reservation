package catalog

import (
	"pension-reservation-api/mod/catalog"

	"gorm.io/gorm"
)

type GetRooms struct {
	db *gorm.DB
}

func NewGetRooms(db *gorm.DB) *GetRooms {
	return &GetRooms{
		db: db,
	}
}

func (m *GetRooms) Run(roomType *int) (catalog.GetRoomsQueryResult, error) {
	r := catalog.GetRoomsQueryResult{}

	return r, nil
}
