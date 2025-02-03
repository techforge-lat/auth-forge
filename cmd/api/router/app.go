package router

import (
	"auth-forge/internal/core/app/infrastructure/in/httprest"
	"auth-forge/pkg/dependency"
	"auth-forge/pkg/server"

	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/linkit"
)

func AppRoutes(server *server.Server) error {
	handler, err := linkit.Resolve[httprest.Handler](server.Container, dependency.AppHandler)
	if err != nil {
		return errortrace.OnError(err)
	}

	group := server.Echo.Group("/api/v1/apps")

	group.POST("", handler.Create)
	group.PUT("/:code", handler.UpdateByCode)
	group.PUT("", handler.Update)
	group.DELETE("/:code", handler.DeleteByCode)
	group.DELETE("", handler.Delete)
	group.GET("/:code", handler.FindOneByCode)
	group.GET("", handler.FindAll)

	return nil
}
