package handlers

import (
	app "main/infrastructure/app"
	"main/infrastructure/http/handlers/health"
)

func SetRoutes(a *app.App) {
	health.RegisterRoutes(a)
	//.RegisterRoutes(a)
}
