package models

type Event struct {
	Id string `json:"__id,omitempty"`

	Title string `json:"summary,omitempty" binding:"required"`

	Description string `json:"description,omitempty" binding:"required"`

	OrganizerName string `json:"organizerName,omitempty" binding:"required"`

	OrganizerEmail string `json:"organizerEmail,omitempty" binding:"required"`

	Location string `json:"location,omitempty" binding:"required"`

	StartTime string `json:"startTime,omitempty" binding:"required"`

	EndTime string `json:"endTime,omitempty" binding:"required"`
}
