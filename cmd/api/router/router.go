package router

import "cloud-crm-backend/pkg/server"

func SetAPIRoutes(server *server.Server) {
	server.Echo.GET("/health", server.HealthCheckController)
}
