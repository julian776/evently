package models

type Event struct {
	Id string `json:"id,omitempty" mapstructure:"id"`

	Title string `json:"title,omitempty" mapstructure:"title"`

	Description string `json:"description,omitempty" mapstructure:"description"`

	Cost float32 `json:"cost,omitempty" mapstructure:"cost"`

	Location string `json:"location,omitempty" mapstructure:"location"`

	Attendees []string `json:"attendees,omitempty" mapstructure:"attendees"`

	OrganizerName string `json:"organizerName,omitempty" mapstructure:"organizerName"`

	OrganizerEmail string `json:"organizerEmail,omitempty" mapstructure:"organizerEmail"`

	StartDate string `json:"startDate,omitempty" mapstructure:"startDate"`

	EndDate string `json:"endDate,omitempty" mapstructure:"endDate"`

	StartTime string `json:"startTime,omitempty" mapstructure:"startTime"`

	EndTime string `json:"endTime,omitempty" mapstructure:"endTime"`
}
