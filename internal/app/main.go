package app

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/gabriwl165/distributed-event-orchestrator/internal/config"
	"github.com/gabriwl165/distributed-event-orchestrator/internal/domain"
	"github.com/gabriwl165/distributed-event-orchestrator/internal/infra/kafka"
	"github.com/gabriwl165/distributed-event-orchestrator/internal/infra/logger"
)

func App(config *config.Config, ctx context.Context) {
	var wg sync.WaitGroup

	topics := strings.Split(config.BrokersList, ",")
	messageChannel := make(chan []byte)
	for _, topic := range topics {
		wg.Add(1)
		go func() {
			defer wg.Done()
			consumer := kafka.NewKafkaConsumer([]string{config.BrokerURL}, "data-orchestrator", topic)
			defer func() {
				err := consumer.Close()
				if err != nil {
					logger.GetLogger().Error("Error while trying to shutdown the consumer for ", topic)
				}
			}()

			consumeMessage(ctx, messageChannel, consumer)
		}()
	}

	go func() {
		for msg := range messageChannel {
			fmt.Printf("Received: %s\n", msg)
		}
	}()

	<-ctx.Done()
	wg.Wait()
}

func consumeMessage(ctx context.Context, messageChannel chan []byte, consumer domain.Consumer) {
	err := consumer.Consume(ctx, func(key, value []byte) error {
		select {
		case messageChannel <- value:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err != nil {
		logger.GetLogger().Error("Consume error: %v", err)
	}
}
