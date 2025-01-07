package di

import (
	"cloud-crm-backend/internal/core/invoiceitem/application"
	"cloud-crm-backend/internal/core/invoiceitem/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/invoiceitem/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideInvoiceItemDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.InvoiceItemRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.InvoiceItemUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.InvoiceItemHandler, handler)
}
