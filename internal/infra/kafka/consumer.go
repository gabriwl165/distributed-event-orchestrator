// internal/infrastructure/kafka/consumer.go
package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, groupID, topic string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			GroupID:  groupID,
			Topic:    topic,
			MinBytes: 1,
			MaxBytes: 10e6,
		}),
	}
}

func (c *KafkaConsumer) Consume(ctx context.Context, handle func(key, value []byte) error) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		if err := handle(msg.Key, msg.Value); err != nil {
			return err
		}
	}
}

func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}
