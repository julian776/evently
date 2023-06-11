package errors

type DuplicateAttendee struct {
	Message string
}

func (e *DuplicateAttendee) Error() string {
	return e.Message
}
