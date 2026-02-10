package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	httpserver "github.com/go-hexagonal-practice/cmd/http_server"
	"github.com/go-hexagonal-practice/internal/adapters/config"
	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre"
	"github.com/go-hexagonal-practice/internal/core/service"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger, client, rdb := httpserver.LoadComponents()

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

	userRepo := postgre.NewUserRepository(client.DB, logger)
	userService := service.NewUserService(userRepo)

	mapBusinessHandler := httpserver.MapBusinessRoutes(logger, rdb, userService)
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
