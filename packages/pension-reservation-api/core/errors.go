package core

type IllegalArguments struct {
	message string
}

var _ error = (*IllegalArguments)(nil)

func (e *IllegalArguments) Error() string {
	return e.message
}
