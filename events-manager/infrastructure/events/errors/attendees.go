package errors

type DuplicateAttendee struct {
	Message string
}

func NewDuplicateAttendeeError() DuplicateAttendee {
	return DuplicateAttendee{
		Message: "Attendee already exists",
	}
}

func (e DuplicateAttendee) Error() string {
	return e.Message
}
