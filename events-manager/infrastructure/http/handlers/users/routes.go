package users

import "events-manager/infrastructure/app"

func RegisterRoutes(a app.App) {
	r := a.Server.Group("/users")

	r.POST("", createUser(*a.CreateUserUseCase))
	r.POST("login", login(*a.LoginUserUseCase))
	r.GET("/:email", getUserByEmail(*a.GetUserByEmailUseCase))
}
