package main

import (
	"log"
	"notifier/infrastructure/rabbit/handlers"
)

func main() {
	app := CreateApp()

	handlers.SetHandlers(app)

	if err := app.Run(); err != nil {
		log.Fatalf("FATAL: %s", err.Error())
	}
}
