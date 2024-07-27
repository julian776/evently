package events

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// For db interactions only
type event struct {
	gorm.Model

	Title string `json:"title,omitempty" binding:"required"`

	Description string `json:"description,omitempty" binding:"required"`

	Cost float32 `json:"cost,omitempty" binding:"gte=0"`

	Location string `json:"location,omitempty" binding:"required"`

	Attendees pq.StringArray `json:"attendees,omitempty" gorm:"type:text[]"`

	OrganizerName string `json:"organizerName,omitempty" binding:"required"`

	OrganizerEmail string `json:"organizerEmail,omitempty" binding:"required"`

	StartDate string `json:"startDate,omitempty" binding:"required"`

	EndDate string `json:"endDate,omitempty" binding:"required"`

	StartTime string `json:"startTime,omitempty" binding:"required"`

	EndTime string `json:"endTime,omitempty" binding:"required"`
}
