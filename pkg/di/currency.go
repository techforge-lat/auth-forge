package di

import (
	"cloud-crm-backend/internal/core/currency/application"
	"cloud-crm-backend/internal/core/currency/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/currency/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideCurrencyDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.CurrencyRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.CurrencyUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.CurrencyHandler, handler)
}
