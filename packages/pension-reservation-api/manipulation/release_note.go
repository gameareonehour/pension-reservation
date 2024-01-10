package manipulation

import (
	release_note_manipulation "pension-reservation-api/manipulation/mod/release_note"
	"pension-reservation-api/mod/release_note"

	"gorm.io/gorm"
)

type ReleaseNoteManipulation struct {
	db *gorm.DB
}

var _ release_note.ReleaseNoteQuery = (*ReleaseNoteManipulation)(nil)

func (m *ReleaseNoteManipulation) GetLatestReleaseNotes() (release_note.GetLatestReleaseNotesQueryResult, error) {
	query := release_note_manipulation.NewGetLatestReleaseNotes(m.db)

	return query.Run()
}

func NewReleaseNoteManipulation(db *gorm.DB) *ReleaseNoteManipulation {
	return &ReleaseNoteManipulation{
		db: db,
	}
}
