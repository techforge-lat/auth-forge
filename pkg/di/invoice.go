package di

import (
	"cloud-crm-backend/internal/core/invoice/application"
	"cloud-crm-backend/internal/core/invoice/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/invoice/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideInvoiceDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.InvoiceRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.InvoiceUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.InvoiceHandler, handler)
}
