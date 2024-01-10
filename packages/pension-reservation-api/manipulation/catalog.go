package manipulation

import (
	catalog_manipulation "pension-reservation-api/manipulation/pkg/catalog"
	"pension-reservation-api/mod/catalog"

	"gorm.io/gorm"
)

type CatalogManipulation struct {
	db *gorm.DB
}

var _ catalog.CatalogQuery = (*CatalogManipulation)(nil)

func NewCatalogManipulation(db *gorm.DB) *CatalogManipulation {
	return &CatalogManipulation{
		db: db,
	}
}

func (m *CatalogManipulation) GetRooms(roomType *int) (catalog.GetRoomsQueryResult, error) {
	query := catalog_manipulation.NewGetRooms(m.db)

	return query.Run(roomType)
}
