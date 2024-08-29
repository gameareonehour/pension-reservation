package catalog

import (
	"pension-reservation-api/mod/catalog"

	"gorm.io/gorm"
)

type queryImpl struct {
	db *gorm.DB
}

var _ catalog.CatalogQuery = (*queryImpl)(nil)

func NewQuery(db *gorm.DB) *queryImpl {
	return &queryImpl{
		db: db,
	}
}

func (q *queryImpl) GetRooms(roomType *int) (*catalog.GetRoomsQueryResult, error) {
	return q.findRooms(roomType)
}
