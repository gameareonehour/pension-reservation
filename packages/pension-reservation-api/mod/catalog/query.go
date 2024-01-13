package catalog

type CatalogQuery interface {
	GetRooms(roomType *int) (*GetRoomsQueryResult, error)
}

type GetRoomsQueryResult []GetRoomsQueryResultInner

type GetRoomsQueryResultInner struct {
	ID       int
	Name     string
	Type     string
	Dayfee   int
	ImageURL string
}
