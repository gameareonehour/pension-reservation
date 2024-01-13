package catalog

import (
	"pension-reservation-api/core"
	"pension-reservation-api/openapi/generated"
	"strconv"
)

type service struct {
	query CatalogQuery
}

func NewService(query CatalogQuery) *service {
	return &service{
		query: query,
	}
}

func (s *service) GetCatalog(params generated.GetCatalogParams) (*generated.GetCatalogResponse, error) {
	response := generated.GetCatalogResponse{}

	roomType, err := s.validateGetCatalogParams(params)
	if err != nil {
		return nil, err
	}

	rooms, err := s.query.GetRooms(roomType)
	if err != nil {
		return nil, err
	}

	for _, v := range *rooms {
		response.Items = append(response.Items, getCatalogResponseItem{
			Id:       v.ID,
			Dayfee:   v.Dayfee,
			ImageUrl: v.ImageURL,
			Name:     v.Name,
			Type:     v.Type,
		})
	}

	return &response, nil
}

func (s *service) validateGetCatalogParams(params generated.GetCatalogParams) (*int, error) {
	roomType := (*int)(nil)

	if params.Type != nil {
		v, err := strconv.Atoi(*params.Type)
		if err != nil {
			return nil, &core.IllegalArguments{}
		}

		roomType = &v
	}

	return roomType, nil
}
