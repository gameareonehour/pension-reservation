package catalog

import (
	"pension-reservation-api/mod/catalog"
	"pension-reservation-api/model"

	"github.com/pkg/errors"
)

func (q *queryImpl) findRooms(roomType *int) (*catalog.GetRoomsQueryResult, error) {
	rooms, err := q.findByRoomType(roomType)
	if err != nil {
		return nil, err
	}

	queryResult := catalog.GetRoomsQueryResult{}

	for _, room := range rooms {
		queryResult = append(queryResult, catalog.GetRoomsQueryResultInner{
			ID:       int(room.ID),
			Name:     room.Name,
			Type:     room.Type.Name,
			Dayfee:   room.Dayfee,
			ImageURL: room.Images[0].ImageURL,
		})
	}

	return &queryResult, nil
}

func (q *queryImpl) findByRoomType(roomType *int) ([]*model.Room, error) {
	rooms := []*model.Room{}
	var err error

	tx := q.db.
		Preload("Type").
		Preload("Images").
		Order("type_id").
		Order("dayfee")

	if roomType != nil {
		err = tx.Find(&rooms, "type_id = ?", roomType).Error
	} else {
		err = tx.Find(&rooms).Error
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return rooms, nil
}
