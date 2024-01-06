package release_note

type GetLatestReleaseNotesQuery interface {
	Run() *GetLatestReleaseNotesQueryResult
}

type GetLatestReleaseNotesQueryResult struct {
	Items []*GetLatestReleaseNotesQueryResultElement
}

type GetLatestReleaseNotesQueryResultElement struct {
	ID        int
	Content   int
	CreatedAt string
	UpdatedAt string
}
