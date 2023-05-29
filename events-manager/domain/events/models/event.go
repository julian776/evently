package models

type Event struct {
	Id string `json:"id,omitempty"`

	Title string `json:"title,omitempty" binding:"required"`

	Description string `json:"description,omitempty" binding:"required"`

	Cost string `json:"cost,omitempty" binding:"required"`

	Location string `json:"location,omitempty" binding:"required"`

	Attendees []string `json:"attendees,omitempty"`

	OrganizerName string `json:"organizerName,omitempty" binding:"required"`

	OrganizerEmail string `json:"organizerEmail,omitempty" binding:"required"`

	StartTime string `json:"startTime,omitempty" binding:"required"`

	EndTime string `json:"endTime,omitempty" binding:"required"`
}
