package router

import (
	"auth-forge/pkg/server"

	"github.com/techforge-lat/errortrace/v2"
)

func SetAPIRoutes(server *server.Server) error {
	server.Echo.GET("/health", server.HealthCheckController)

	if err := AppRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := TenantRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}
