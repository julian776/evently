package main

import (
	"log"
)

func main() {
	app := CreateApp()

	if err := app.Run(); err != nil {
		log.Fatalf("FATAL: %s", err.Error())
	}
}
