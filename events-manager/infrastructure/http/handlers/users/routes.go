package users

import "events-manager/infrastructure/app"

func RegisterRoutes(a *app.App) {
	r := a.Server.Group("/users")

	r.POST("", createUser(*a.CreateUserUseCase))
	// r.GET("/", getAllEvents(*a.GetAllEventsUseCase))
	// r.GET("/:id", getEventById(*a.GetEventByIdUseCase))
	// r.DELETE("/:id", deleteEventById(*a.DeleteEventByIdUseCase))
	// r.PUT("/", updateEvent(*a.UpdateEventUseCase))
}
