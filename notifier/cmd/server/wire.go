//go:build wireinject
// +build wireinject

package main

import (
	"notifier/domain/broker"
	remindersD "notifier/domain/reminders/repositories"
	"notifier/infrastructure/app"
	"notifier/infrastructure/rabbit"
	"notifier/infrastructure/reminders"
	"notifier/pkgs/logger"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func CreateApp() *app.App {
	wire.Build(
		app.SettingsProvider,
		reminders.NewRemindersMongoRepository,
		wire.Bind(new(remindersD.RemindersRepository), new(*reminders.RemindersMongoRepository)),
		logger.NewLogger,
		wire.Bind(new(logger.Logger), new(*zap.SugaredLogger)),
		rabbit.NewRabbitListener,
		wire.Bind(new(broker.BrokerListener), new(*rabbit.RabbitListener)),
		app.NewApp,
	)
	return new(app.App)
}
