package events

import (
	"events-manager/domain/events/models"
	"fmt"
)

func mapEventToPostgresEvent(e models.Event) *event {
	return &event{
		Title:          e.Title,
		Description:    e.Description,
		Cost:           e.Cost,
		Location:       e.Location,
		Attendees:      e.Attendees,
		OrganizerName:  e.OrganizerName,
		OrganizerEmail: e.OrganizerEmail,
		StartDate:      e.StartDate,
		EndDate:        e.EndDate,
		StartTime:      e.StartTime,
		EndTime:        e.EndTime,
	}
}

func mapPostgresEventToEvent(e event) *models.Event {
	return &models.Event{
		Id:             fmt.Sprint(e.ID),
		Title:          e.Title,
		Description:    e.Description,
		Cost:           e.Cost,
		Location:       e.Location,
		Attendees:      e.Attendees,
		OrganizerName:  e.OrganizerName,
		OrganizerEmail: e.OrganizerEmail,
		StartDate:      e.StartDate,
		EndDate:        e.EndDate,
		StartTime:      e.StartTime,
		EndTime:        e.EndTime,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}
