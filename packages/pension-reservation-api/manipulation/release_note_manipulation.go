package manipulation

import (
	release_note_manipulation "pension-reservation-api/manipulation/mod/release_note_manipulation"
	"pension-reservation-api/mod/release_note"

	"gorm.io/gorm"
)

type releaseNoteManipulation struct {
	db *gorm.DB
}

var _ release_note.ReleaseNoteQuery = (*releaseNoteManipulation)(nil)

func (m *releaseNoteManipulation) GetLatestReleaseNotes() (release_note.GetLatestReleaseNotesQueryResult, error) {
	query := release_note_manipulation.NewGetLatestReleaseNotes(m.db)

	return query.Run()
}

func NewReleaseNoteManipulation(db *gorm.DB) *releaseNoteManipulation {
	return &releaseNoteManipulation{
		db: db,
	}
}
