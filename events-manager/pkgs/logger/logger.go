package logger

import (
	"log"
	"strings"

	"go.uber.org/zap"
)

func NewLogger(settings Settings) *zap.SugaredLogger {
	var zapLogger *zap.Logger
	var err error

	if strings.ToUpper(settings.Enviroment) == "DEV" {
		zapLogger, err = zap.NewDevelopment()
	} else {
		zapLogger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatalf("can't initialize logger: %v", err)
	}

	return zapLogger.Sugar()
}

type Logger interface {
	Debugf(msg string, keysAndValues ...interface{})
	Warnf(msg string, keysAndValues ...interface{})
	Errorf(msg string, keysAndValues ...interface{})
	Fatalf(msg string, keysAndValues ...interface{})
	Sync() error
}
