package router

import (
	"cloud-crm-backend/pkg/server"

	"github.com/techforge-lat/errortrace/v2"
)

func SetAPIRoutes(server *server.Server) error {
	server.Echo.GET("/health", server.HealthCheckController)

	if err := tenantRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}
