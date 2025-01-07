package di

import (
	"cloud-crm-backend/internal/core/supplier/application"
	"cloud-crm-backend/internal/core/supplier/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/supplier/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideSupplierDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.SupplierRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.SupplierUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.SupplierHandler, handler)
}
