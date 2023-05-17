package handlers

import (
	"events-manager/infrastructure/app"
	"events-manager/infrastructure/http/handlers/health"
)

func SetRoutes(a *app.App) {
	health.RegisterRoutes(a)
	//.RegisterRoutes(a)
}
