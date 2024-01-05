package manipulation

import (
	"pension-reservation-api/mod/release_note"

	"gorm.io/gorm"
)

type GetLatestReleaseNotes struct {
	db *gorm.DB
}

var _ release_note.GetLatestReleaseNotesQuery = (*GetLatestReleaseNotes)(nil)

func NewGetLatestReleaseNotes(db *gorm.DB) *GetLatestReleaseNotes {
	return &GetLatestReleaseNotes {
		db: db,
	}
}

func (m *GetLatestReleaseNotes) Run() *release_note.GetLatestReleaseNotesQueryResult {
	return &release_note.GetLatestReleaseNotesQueryResult{}
}
