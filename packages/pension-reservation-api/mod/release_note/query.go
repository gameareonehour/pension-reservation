package release_note

import "time"

type ReleaseNoteQuery interface {
	GetLatestReleaseNotes() (GetLatestReleaseNotesQueryResult, error)
}

type GetLatestReleaseNotesQueryResult []GetLatestReleaseNotesQueryResultInner

type GetLatestReleaseNotesQueryResultInner struct {
	ID        int
	Text      string
	CreatedAt time.Time
}
