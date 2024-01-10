package manipulation

import (
	catalog_manipulation "pension-reservation-api/manipulation/mod/catalog"
	"pension-reservation-api/mod/catalog"

	"gorm.io/gorm"
)

type catalogManipulation struct {
	db *gorm.DB
}

var _ catalog.CatalogQuery = (*catalogManipulation)(nil)

func NewCatalogManipulation(db *gorm.DB) *catalogManipulation {
	return &catalogManipulation{
		db: db,
	}
}

func (m *catalogManipulation) GetRooms(roomType *int) (catalog.GetRoomsQueryResult, error) {
	query := catalog_manipulation.NewGetRooms(m.db)

	return query.Run(roomType)
}
