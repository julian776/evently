package models

type Event struct {
	Id string `json:"__id,omitempty"`

	Title string `json:"summary,omitempty"`

	Description string `json:"description,omitempty"`

	UserId string `json:"userId,omitempty"`

	Location string `json:"location,omitempty"`

	StartTime string `json:"start_time,omitempty"`

	EndTime string `json:"end_time,omitempty"`

	Attendess []string `json:"attendess,omitempty"`

	// Set reminders to a given event.
	// Check valid reminders.
	Reminders map[string]bool `json:"reminders,omitempty"`
}
