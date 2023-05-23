package main

import (
	"events-manager/infrastructure/events"
	"events-manager/infrastructure/http/handlers"
	"log"
)

func main() {
	app := CreateApp()

	events.LoadEventsService(app.BrokerPublisher, *app.Settings.EventsSettings)

	handlers.SetRoutes(app)

	if err := app.Run(); err != nil {
		log.Fatalf("FATAL: %s", err.Error())
	}
}
