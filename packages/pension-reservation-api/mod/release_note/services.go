package release_note

import (
	"pension-reservation-api/openapi/generated"
	"time"
)

type GetLatestReleaseNotesService struct {
	query ReleaseNoteQuery
}

func NewGetLatestReleaseNotesService(query ReleaseNoteQuery) *GetLatestReleaseNotesService {
	return &GetLatestReleaseNotesService{
		query: query,
	}
}

func (s *GetLatestReleaseNotesService) GetLatestReleaseNotes() (*generated.GetReleaseNotesResponse, error) {
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
