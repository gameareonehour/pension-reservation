package release_note

import (
	"pension-reservation-api/openapi/generated"
	"time"
)

type service struct {
	query ReleaseNoteQuery
}

func NewService(query ReleaseNoteQuery) *service {
	return &service{
		query: query,
	}
}

func (s *service) GetLatestReleaseNotes() (*generated.GetReleaseNotesResponse, error) {
	response := generated.GetReleaseNotesResponse{}
	
	notes, err := s.query.GetLatestReleaseNotes()
	if err != nil {
		return nil, err
	}

	for _, v := range notes {
		response.Items = append(response.Items, GetReleaseNotesResponseItem{
			Id:        v.ID,
			Text:      v.Text,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
		})
	}

	return &response, nil
}
