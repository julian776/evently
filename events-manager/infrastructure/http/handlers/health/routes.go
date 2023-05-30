package health

import "events-manager/infrastructure/app"

func RegisterRoutes(a *app.App) {
	a.Server.GET("/health", healthCheck())
}
