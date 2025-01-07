package di

import (
	"cloud-crm-backend/internal/core/contract/application"
	"cloud-crm-backend/internal/core/contract/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/contract/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideContractDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.ContractRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.ContractUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.ContractHandler, handler)
}
