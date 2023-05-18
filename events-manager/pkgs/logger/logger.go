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
	Debugf(template string, keysAndValues ...interface{})

	Warnf(template string, keysAndValues ...interface{})

	Errorf(template string, keysAndValues ...interface{})

	Fatalf(template string, keysAndValues ...interface{})

	// Sync flushes any buffered log entries.
	Sync() error
}
