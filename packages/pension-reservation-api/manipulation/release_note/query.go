package release_note

import (
	"pension-reservation-api/mod/release_note"

	"gorm.io/gorm"
)

type queryImpl struct {
	db *gorm.DB
}

var _ release_note.ReleaseNoteQuery = (*queryImpl)(nil)

func NewQuery(db *gorm.DB) *queryImpl {
	return &queryImpl{
		db: db,
	}
}

func (q *queryImpl) GetLatestReleaseNotes() (release_note.GetLatestReleaseNotesQueryResult, error) {
	return q.findLatestReleaseNotes()
}
