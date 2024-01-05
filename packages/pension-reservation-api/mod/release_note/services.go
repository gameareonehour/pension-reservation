package release_note

type GetLatestReleaseNotesService struct {
	query GetLatestReleaseNotesQuery
}

func NewGetLatestReleaseNotesService(query GetLatestReleaseNotesQuery) *GetLatestReleaseNotesService {
	return &GetLatestReleaseNotesService{
		query: query,
	}
}

func (h *GetLatestReleaseNotesService) GetLatestReleaseNotes() *GetLatestReleaseNotesQueryResult {
	return h.query.Run()
}
