package di

import (
	"cloud-crm-backend/internal/core/contractproduct/application"
	"cloud-crm-backend/internal/core/contractproduct/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/contractproduct/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideContractProductDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.ContractProductRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.ContractProductUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.ContractProductHandler, handler)
}
