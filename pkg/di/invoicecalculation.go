package di

import (
	"cloud-crm-backend/internal/core/invoicecalculation/application"
	"cloud-crm-backend/internal/core/invoicecalculation/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/invoicecalculation/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideInvoiceCalculationDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.InvoiceCalculationRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.InvoiceCalculationUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.InvoiceCalculationHandler, handler)
}
