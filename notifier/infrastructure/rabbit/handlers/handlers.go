package handlers

import (
	"notifier/infrastructure/app"
	"notifier/infrastructure/rabbit/handlers/events"
)

func SetHandlers(a *app.App) {
	events.SetHandlers(a)
}
