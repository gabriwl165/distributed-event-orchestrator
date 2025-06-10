package kafka_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWithKafka(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "confluentinc/cp-kafka:latest",
		ExposedPorts: []string{"9092/tcp", "29092/tcp"},
		Env: map[string]string{
			"KAFKA_BROKER_ID":                        "1",
			"KAFKA_ZOOKEEPER_CONNECT":                "localhost:2181",
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP":   "PLAINTEXT:PLAINTEXT,PLAINTEXT_INTERNAL:PLAINTEXT",
			"KAFKA_ADVERTISED_LISTENERS":             "PLAINTEXT://localhost:9092,PLAINTEXT_INTERNAL://localhost:29092",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR": "1",
			"KAFKA_LISTENERS":                        "PLAINTEXT://0.0.0.0:9092,PLAINTEXT_INTERNAL://0.0.0.0:29092",
			"KAFKA_INTER_BROKER_LISTENER_NAME":       "PLAINTEXT_INTERNAL",
		},
		WaitingFor: wait.ForListeningPort("9092/tcp"),
	}

	kafkaC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = kafkaC.Terminate(ctx)
	})

	host, err := kafkaC.Host(ctx)
	require.NoError(t, err)

	port, err := kafkaC.MappedPort(ctx, "29092")
	require.NoError(t, err)

	broker := fmt.Sprintf("%s:%s", host, port.Port())
	require.NotEmpty(t, broker)
}
