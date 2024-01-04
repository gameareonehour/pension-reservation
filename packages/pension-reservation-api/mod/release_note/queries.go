package release_note

type GetLatestReleaseNotesQuery interface {
	Run() queryResult
}

type queryResult interface {
	Items() []*queryResultElement
}

type queryResultElement interface {
	ID() int
	Content() int
	CreatedAt() string
	UpdatedAt() string
}
