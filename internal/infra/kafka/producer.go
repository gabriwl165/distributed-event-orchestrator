package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string) *KafkaPublisher {
	return &KafkaPublisher{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *KafkaPublisher) Produce(ctx context.Context, topic string, key, value []byte) error {
	msg := kafka.Message{
		Topic: topic,
		Key:   key,
		Value: value,
	}
	return p.writer.WriteMessages(ctx, msg)
}
