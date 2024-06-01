package catalog

import (
	"pension-reservation-api/mod/catalog"

	"gorm.io/gorm"
)

type catalogQueryImpl struct {
	db *gorm.DB
}

var _ catalog.CatalogQuery = (*catalogQueryImpl)(nil)

func NewQuery(db *gorm.DB) *catalogQueryImpl {
	return &catalogQueryImpl{
		db: db,
	}
}

func (q *catalogQueryImpl) GetRooms(roomType *int) (*catalog.GetRoomsQueryResult, error) {
	query := NewGetRooms(q.db)

	return query.Run(roomType)
}
