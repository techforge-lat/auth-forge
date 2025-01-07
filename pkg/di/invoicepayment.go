package di

import (
	"cloud-crm-backend/internal/core/invoicepayment/application"
	"cloud-crm-backend/internal/core/invoicepayment/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/invoicepayment/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideInvoicePaymentDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.InvoicePaymentRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.InvoicePaymentUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.InvoicePaymentHandler, handler)
}
