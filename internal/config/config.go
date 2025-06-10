package config

type Config struct {
	BrokersList string `yaml:"brokers"`
	BrokerPort  string `yaml:"brokerUrl"`
	BrokerURL   string `yaml:"brokerPort"`
}
