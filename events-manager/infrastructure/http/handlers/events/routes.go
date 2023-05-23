package events

import "events-manager/infrastructure/app"

func RegisterRoutes(a *app.App) *app.App {
	r := a.Server.Group("/events")

	r.POST("/", createEvent(*a.CreateEventUseCase))

	return a
}
