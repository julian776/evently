package events

import "events-manager/infrastructure/app"

func RegisterRoutes(a *app.App) *app.App {
	r := a.Server.Group("/events")

	r.POST("/", createEvent(*a.CreateEventUseCase))
	r.GET("/:id", getEventById(*a.GetEventByIdUseCase))
	r.DELETE("/:id", deleteEventById(*a.DeleteEventByIdUseCase))
	r.PUT("/", updateEvent(*a.UpdateEventUseCase))
	return a
}
