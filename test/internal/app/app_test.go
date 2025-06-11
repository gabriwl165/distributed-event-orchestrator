package app

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/gabriwl165/distributed-event-orchestrator/internal/app"
	"github.com/gabriwl165/distributed-event-orchestrator/internal/config"
	"github.com/gabriwl165/distributed-event-orchestrator/internal/infra/kafka"
	"github.com/gabriwl165/distributed-event-orchestrator/internal/infra/logger"
	"github.com/gabriwl165/distributed-event-orchestrator/test/mock"
)

func mockConfig(brokerURL string, brokerList string) *config.Config {
	config := &config.Config{
		BrokersList: brokerList,
		BrokerURL:   brokerURL,
	}
	return config
}

func TestAppKafkaContainer(t *testing.T) {
	logger := logger.GetLogger()
	ctx, cancel := context.WithCancel(context.Background())

	kafkaContainer, brokerUrl, err := mock.StartKafkaContainer(ctx)
	if err != nil {
		logger.Fatalf("Kafka container error: %s", err)
	}

	logger.Info("Kafka broker running at:", brokerUrl)

	brokers := "order-create-action,order-update-action"

	go func() {
		data := map[string]interface{}{
			"name": "Gabriel",
			"age":  29,
		}

		value, _ := json.Marshal(data)

		producer := kafka.NewKafkaProducer([]string{brokerUrl})
		if err := producer.Produce(ctx, "order-create-action", []byte("order-create-action"), value); err != nil {
			logger.Errorf("Error producing message: %v", err)
		}
		time.Sleep(5 * time.Second)
		cancel()

		if err := kafkaContainer.Terminate(ctx); err != nil {
			logger.Errorf("Error shutting down the container: %v", err)
		}

	}()

	config := mockConfig(brokerUrl, brokers)
	app.App(config, ctx)

	print("ufa")
}
