package handlers

import (
	"events-manager/infrastructure/app"
	"events-manager/infrastructure/http/handlers/events"
	"events-manager/infrastructure/http/handlers/health"
	"events-manager/infrastructure/http/handlers/users"
)

func SetRoutes(a *app.App) {
	health.RegisterRoutes(a)
	events.RegisterRoutes(a)
	users.RegisterRoutes(a)
}
