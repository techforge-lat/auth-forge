package di

import (
	"cloud-crm-backend/internal/core/invoicecalculationitem/application"
	"cloud-crm-backend/internal/core/invoicecalculationitem/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/invoicecalculationitem/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideInvoiceCalculationItemDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.InvoiceCalculationItemRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.InvoiceCalculationItemUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.InvoiceCalculationItemHandler, handler)
}
