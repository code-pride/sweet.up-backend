package util

import (
	"log"

	"go.uber.org/zap"
)

type LoggerConfiguration struct {
	RootLevel string
}

func ConfigureLogger(logConf LoggerConfiguration) *zap.SugaredLogger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	return logger.Sugar()
}
