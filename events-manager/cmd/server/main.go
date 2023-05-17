package main

import (
	"context"
	"log"
	"main/infrastructure/http/handlers"
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
