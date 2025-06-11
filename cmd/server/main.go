package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gabriwl165/distributed-event-orchestrator/internal/app"
	"github.com/gabriwl165/distributed-event-orchestrator/internal/config"
	"github.com/gabriwl165/distributed-event-orchestrator/internal/infra/logger"
	"gopkg.in/yaml.v3"
)

func main() {
	log.Print("Olá Mundo")
	config, _ := loadConfig()
	logger := logger.GetLogger()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	if config == nil {
		logger.Error("Failed to load Application Config")
	}

	logger.Info("Orchestrator Running")

	go func() {
		<-ctx.Done()
		logger.Info("Shutting down gracefully...")
		time.Sleep(1 * time.Second) // simulate cleanup
		os.Exit(0)
	}()
	app.App(config, ctx)
}

func loadConfig() (*config.Config, error) {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}

	var cfg config.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil

}
