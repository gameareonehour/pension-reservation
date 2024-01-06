package release_note

type GetLatestReleaseNotesQuery interface {
	Run() (GetLatestReleaseNotesQueryResult, error)
}

type GetLatestReleaseNotesQueryResult []GetLatestReleaseNotesQueryResultInner

type GetLatestReleaseNotesQueryResultInner struct {
	ID        int   
	Text      string
	CreatedAt string
	UpdatedAt string
}
