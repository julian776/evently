package health

import "main/infrastructure/app"

func RegisterRoutes(a *app.App) *app.App {
	a.Server.GET("/health", healthCheck())

	return a
}
