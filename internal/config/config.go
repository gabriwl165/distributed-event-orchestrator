package config

type Config struct {
	BrokersList string `yaml:"brokers"`
	BrokerURL   string `yaml:"brokerPort"`
}
