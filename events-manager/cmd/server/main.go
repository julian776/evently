package main

import (
	"context"
	"events-manager/infrastructure/http/handlers"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	app := CreateApp()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	cancel()

	handlers.SetRoutes(app)

	if err := app.Run(ctx); err != nil {
		log.Fatalf("FATAL: %s", err.Error())
	}
}
