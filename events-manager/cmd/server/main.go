package main

import (
	"events-manager/infrastructure/http/handlers"
	"log"
)

func main() {
	app := CreateApp()

	handlers.SetRoutes(*app)

	if err := app.Run(); err != nil {
		log.Fatalf("FATAL: %s", err.Error())
	}
}
