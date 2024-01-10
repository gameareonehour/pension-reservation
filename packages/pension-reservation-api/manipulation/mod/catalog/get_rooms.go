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
	// rooms := []model.Room{}

	// if roomType != nil {

	// }
	// m.db.Where()	

	r := catalog.GetRoomsQueryResult{}

	return r, nil
}

// func (m *GetRooms) getRoomsByType(roomType *int) error {

// 	return nil
// 	// m.db.Where(&model.Room{ TypeID:  })
// }
