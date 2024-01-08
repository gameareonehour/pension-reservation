package release_note

import "time"

type GetLatestReleaseNotesQuery interface {
	Run() (GetLatestReleaseNotesQueryResult, error)
}

type GetLatestReleaseNotesQueryResult []GetLatestReleaseNotesQueryResultInner

type GetLatestReleaseNotesQueryResultInner struct {
	ID        int
	Text      string
	CreatedAt time.Time
}
