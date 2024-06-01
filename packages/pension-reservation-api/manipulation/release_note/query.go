package release_note

import (
	"pension-reservation-api/mod/release_note"

	"gorm.io/gorm"
)

type releaseNoteQueryImpl struct {
	db *gorm.DB
}

var _ release_note.ReleaseNoteQuery = (*releaseNoteQueryImpl)(nil)

func NewQuery(db *gorm.DB) *releaseNoteQueryImpl {
	return &releaseNoteQueryImpl{
		db: db,
	}
}

func (q *releaseNoteQueryImpl) GetLatestReleaseNotes() (release_note.GetLatestReleaseNotesQueryResult, error) {
	query := NewGetLatestReleaseNotes(q.db)

	return query.Run()
}
