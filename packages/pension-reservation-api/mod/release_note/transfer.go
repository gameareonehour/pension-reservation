package release_note

type GetReleaseNotesResponseItem struct {
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	Text      string `json:"text"`
	UpdatedAt string `json:"updated_at"`
}
