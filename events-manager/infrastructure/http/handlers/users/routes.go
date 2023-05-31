package users

import "events-manager/infrastructure/app"

func RegisterRoutes(a app.App) {
	r := a.Server.Group("/users")

	r.POST("", createUser(*a.CreateUserUseCase))
	r.GET("/:email", getUserByEmail(*a.GetUserByEmailUseCase))
	// r.GET("/", getAllEvents(*a.GetAllEventsUseCase))
	// r.DELETE("/:id", deleteEventById(*a.DeleteEventByIdUseCase))
	// r.PUT("/", updateEvent(*a.UpdateEventUseCase))
}
