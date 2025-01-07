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

	if err := supplierRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := ProductRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := ProductPriceRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := ContractRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := ContractProductRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := InvoiceRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := InvoiceItemRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := InvoicePaymentRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := InvoiceCalculationRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := InvoiceCalculationItemRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := CurrencyRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	if err := PaymentAccountRoutes(server); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}
