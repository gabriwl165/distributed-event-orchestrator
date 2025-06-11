package mock

import (
	"context"
	"fmt"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartKafkaContainer(ctx context.Context) (testcontainers.Container, string, error) {
	zookeeperReq := testcontainers.ContainerRequest{
		Image:        "confluentinc/cp-zookeeper:latest",
		ExposedPorts: []string{"2181/tcp"},
		Env: map[string]string{
			"ZOOKEEPER_CLIENT_PORT": "2181",
			"ZOOKEEPER_TICK_TIME":   "2000",
		},
		WaitingFor: wait.ForListeningPort("2181/tcp"),
	}
	zookeeper, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: zookeeperReq,
		Started:          true,
	})
	if err != nil {
		return nil, "", err
	}

	zkHost, err := zookeeper.Host(ctx)
	if err != nil {
		return nil, "", err
	}
	zkPort, err := zookeeper.MappedPort(ctx, "2181")
	if err != nil {
		return nil, "", err
	}
	zkAddress := fmt.Sprintf("%s:%s", zkHost, zkPort.Port())

	kafkaReq := testcontainers.ContainerRequest{
		Image:        "confluentinc/cp-kafka:latest",
		ExposedPorts: []string{"9092/tcp"},
		Env: map[string]string{
			"KAFKA_BROKER_ID":                        "1",
			"KAFKA_ZOOKEEPER_CONNECT":                zkAddress,
			"KAFKA_ADVERTISED_LISTENERS":             "PLAINTEXT://localhost:9092",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR": "1",
			"KAFKA_LISTENERS":                        "PLAINTEXT://0.0.0.0:9092",
			"KAFKA_INTER_BROKER_LISTENER_NAME":       "PLAINTEXT",
		},
		WaitingFor: wait.ForListeningPort("9092/tcp").WithStartupTimeout(2 * time.Minute),
	}
	kafka, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: kafkaReq,
		Started:          true,
	})
	if err != nil {
		return nil, "", err
	}

	host, _ := kafka.Host(ctx)
	port, _ := kafka.MappedPort(ctx, "9092")
	broker := fmt.Sprintf("%s:%s", host, port.Port())

	return kafka, broker, nil
}
