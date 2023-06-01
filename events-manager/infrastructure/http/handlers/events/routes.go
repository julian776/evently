package events

import "events-manager/infrastructure/app"

func RegisterRoutes(a app.App) {
	r := a.Server.Group("/events")

	r.POST("", createEvent(*a.CreateEventUseCase))
	r.GET("", getAllEvents(*a.GetAllEventsUseCase))
	r.GET("/:id", getEventById(*a.GetEventByIdUseCase))
	r.PUT("", updateEvent(*a.UpdateEventUseCase))
	r.PUT("/attendees", addAttendeeEvent(*a.AddAttendeeEventUseCase))
	r.DELETE("/:id", deleteEventById(*a.DeleteEventByIdUseCase))
}
