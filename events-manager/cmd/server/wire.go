//go:build wireinject
// +build wireinject

package main

import (
	"events-manager/domain/broker"
	"events-manager/domain/events/repositories"
	events "events-manager/domain/events/usecases"
	"events-manager/infrastructure/app"
	postgredb "events-manager/infrastructure/events/adapters/postgre_db"
	client "events-manager/infrastructure/http/client"
	server "events-manager/infrastructure/http/server"
	"events-manager/infrastructure/rabbit"
	"events-manager/pkgs/logger"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func CreateApp() *app.App {
	wire.Build(
		app.SettingsProvider,
		events.UseCasesProvider,
		postgredb.NewPostgreRepository,
		wire.Bind(new(repositories.EventsRepository), new(*postgredb.PostgreRepository)),
		logger.NewLogger,
		wire.Bind(new(logger.Logger), new(*zap.SugaredLogger)),
		server.NewServer,
		client.NewClient,
		rabbit.NewRabbitPublisher,
		wire.Bind(new(broker.BrokerPublisher), new(*rabbit.RabbitPublisher)),
		app.NewApp,
	)
	return new(app.App)
}
