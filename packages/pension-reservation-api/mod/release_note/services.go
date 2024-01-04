package release_note

type GetLatestReleaseNotesService struct {
	query GetLatestReleaseNotesQuery
}

func NewGetLatestReleaseNotesService(query GetLatestReleaseNotesQuery) *GetLatestReleaseNotesService {
	return &GetLatestReleaseNotesService{
		query: query,
	}
}

func (h *GetLatestReleaseNotesService) GetLatestReleaseNotes() queryResult {
	return h.query.Run()
}
