package router

import (
	"auth-forge/internal/core/tenant/infrastructure/in/httprest"
	"auth-forge/pkg/dependency"
	"auth-forge/pkg/server"

	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/linkit"
)

func TenantRoutes(server *server.Server) error {
	handler, err := linkit.Resolve[httprest.Handler](server.Container, dependency.TenantHandler)
	if err != nil {
		return errortrace.OnError(err)
	}

	group := server.Echo.Group("/api/v1/tenants")

	group.POST("", handler.Create)
	group.PUT("", handler.Update)
	group.PUT("/:code", handler.UpdateByCode)
	group.DELETE("/:code", handler.DeleteByCode)
	group.DELETE("", handler.Delete)
	group.GET("/:code", handler.FindOneByCode)
	group.GET("", handler.FindAll)

	return nil
}
