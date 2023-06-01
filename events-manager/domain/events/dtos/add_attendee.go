package dtos

type AddAttendeDTO struct {
	EventId       string `json:"eventId,omitempty" binding:"required"`
	AttendeeEmail string `json:"attendeeEmail,omitempty" binding:"required"`
}
