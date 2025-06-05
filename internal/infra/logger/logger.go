package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	once   sync.Once
	logger *zap.SugaredLogger
)

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		zapLogger, _ := zap.NewProduction()
		logger = zapLogger.Sugar()
	})
	return logger
}
