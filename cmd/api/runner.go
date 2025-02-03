package main

import (
	"auth-forge/cmd"
	"auth-forge/cmd/api/router"
	"auth-forge/internal/shared/domain"
	"auth-forge/pkg/di"
	"auth-forge/pkg/logger"
	"context"
	"os"
)

func Run() {
	logger := logger.NewZeroLog(serviceName)
	server, err := cmd.NewServerInstance(os.Getenv("CONFIGURATION_FILEPATH"), serviceName)
	if err != nil {
		logger.Error(context.Background(), err.Error())
		return
	}

	if err := di.ProvideDependencies(server.Container); err != nil {
		logger.Error(context.Background(), err.Error())
		return
	}

	if err := router.SetAPIRoutes(server); err != nil {
		logger.Error(context.Background(), "failed to register routes", "error", err.Error())
		return
	}

	if err := server.Start(); err != nil {
		logger.Error(context.Background(), err.Error())
	}
}

// RunWithCustomConfig is a function that receives a custom configuration and returns a server instance
// it is useful for testing purposes
func RunWithCustomConfig(config domain.Configuration) {
	logger := logger.NewZeroLog(serviceName)

	server, err := cmd.NewServerInstanceWithCustomConfig(config, serviceName)
	if err != nil {
		logger.Error(context.Background(), err.Error())
		return
	}

	if err := di.ProvideDependencies(server.Container); err != nil {
		logger.Error(context.Background(), err.Error())
		return
	}

	if err := router.SetAPIRoutes(server); err != nil {
		logger.Error(context.Background(), "failed to register routes", "error", err.Error())
		return
	}

	if err := server.Start(); err != nil {
		logger.Error(context.Background(), err.Error())
	}
}
