package logger_test

import (
	"testing"

	"github.com/gabriwl165/distributed-event-orchestrator/internal/infra/logger"
	"github.com/stretchr/testify/assert"
)

func TestGetLoggerSingleton(t *testing.T) {
	logger1 := logger.GetLogger()
	logger2 := logger.GetLogger()

	assert.Same(t, logger1, logger2, "GetLogger should return the same instance")
}
