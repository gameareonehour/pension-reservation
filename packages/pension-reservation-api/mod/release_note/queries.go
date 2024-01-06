package release_note

type GetLatestReleaseNotesQuery interface {
	Run() *GetLatestReleaseNotesQueryResult
}

type GetLatestReleaseNotesQueryResult struct {
	Items []*GetLatestReleaseNotesQueryResultElement
}

type GetLatestReleaseNotesQueryResultElement struct {
	ID        int    
	Text      string
	CreatedAt string
	UpdatedAt string
}
