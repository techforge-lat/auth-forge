package cmd

import (
	"cloud-crm-backend/internal/shared/domain"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"
	"cloud-crm-backend/pkg/localconfig"
	"cloud-crm-backend/pkg/logger"
	"cloud-crm-backend/pkg/server"

	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/linkit"
)

// NewServerInstance is a function that receives a configuration path and a service name and returns a server instance.
func NewServerInstance(configPath, serviceName string) (*server.Server, error) {
	config, err := localconfig.NewLocalConfig(configPath)
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	return NewServerInstanceWithCustomConfig(config, serviceName)
}

// NewServerInstanceWithCustomConfig is a function that receives a custom configuration and a service name and returns a server instance
// it is useful for testing purposes.
func NewServerInstanceWithCustomConfig(config domain.Configuration, serviceName string) (*server.Server, error) {
	container := linkit.New()
	container.Provide(dependency.LocalConfig, config)

	logger := logger.NewZeroLog(serviceName)
	container.Provide(dependency.Logger, logger)

	db, err := database.New(config)
	if err != nil {
		return nil, errortrace.OnError(err)
	}
	container.Provide(dependency.Database, db)

	server, err := server.New(container, serviceName)
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	return server, nil
}
