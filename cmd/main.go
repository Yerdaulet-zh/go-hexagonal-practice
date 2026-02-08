package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	httpserver "github.com/go-hexagonal-practice/cmd/http_server"
	"github.com/go-hexagonal-practice/internal/adapters/config"
	"github.com/go-hexagonal-practice/internal/adapters/logging"
	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre"
	"github.com/go-hexagonal-practice/internal/core/ports"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger, client, rdb := loadComponents()

	defer func() {
		logger.Info("Closing infrastructure connections...")
		if err := client.Close(); err != nil {
			logger.Error("Postgre close error", "error", err)
		}
		if err := rdb.Close(); err != nil {
			logger.Error("Redis close error", "error", err)
		}
		logger.Info("Done")
	}()

	logger.Info("Loading HTTP Server config")
	httpConfig := config.NewHttpConfig()
	logger.Info("Successfully loaded HTTP Server config")

	reg := prometheus.NewRegistry()
	mapBusinessHandler := httpserver.MapBusinessRoutes(logger, rdb)
	mapManagementRoutes := httpserver.MapManagementRoutes(logger, client, reg)

	go func() {
		if err := httpserver.Run(ctx, logger, mapManagementRoutes, httpConfig.HttpManagementAddr(), "Management"); err != nil {
			logger.Error("HTTP Management server error while shutting down", "error", err)
		}
	}()

	if err := httpserver.Run(ctx, logger, mapBusinessHandler, httpConfig.HttpBusinessAddr(), "Business"); err != nil {
		logger.Error("HTTP Business server error while shutting down", "error", err)
		os.Exit(1)
	}

	logger.Info("Application exited cleanly")
}

func loadComponents() (ports.Logger, *postgre.Client, *redis.Client) {
	// Configuration
	cfg, err := config.NewLoggingConfig()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Logger
	logger := logging.NewLogger(cfg)
	logger.Info("Logging successfully configured to use the adapter: ", cfg.Adapter())

	// PostgreSQL
	logger.Info("Loading PostgreSQL config")
	postgreConfig, err := config.NewDefaultDBConfig()
	if err != nil {
		logger.Error("Failed to load PostgreSQL config", "error", err)
		os.Exit(1)
	}

	logger.Info("Connecting to PostgreSQL database")
	client, err := postgre.NewPostgreSQLClient(postgreConfig)
	if err != nil {
		logger.Error("Postgresql connection error", "error", err)
		os.Exit(1)
	}
	logger.Info("Successful PostgreSQL connection")

	// Redis
	logger.Info("Connecting to redis server")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	logger.Info("Successful redis connection")

	return logger, client, rdb
}
