package router

import (
	"cloud-crm-backend/internal/core/supplier/infrastructure/in/httprest"
	"cloud-crm-backend/pkg/dependency"
	"cloud-crm-backend/pkg/server"

	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/linkit"
)

func supplierRoutes(server *server.Server) error {
	handler, err := linkit.Resolve[httprest.Handler](server.Container, dependency.SupplierHandler)
	if err != nil {
		return errortrace.OnError(err)
	}

	group := server.Echo.Group("v1/suppliers")

	group.POST("", handler.Create)
	group.PUT("", handler.Update)
	group.PUT("/:id", handler.UpdateByID)
	group.DELETE("/:id", handler.DeleteByID)
	group.DELETE("", handler.Delete)
	group.GET("/:id", handler.FindOneByID)
	group.GET("", handler.FindOne)
	group.GET("/all", handler.FindAll)

	return nil
}
