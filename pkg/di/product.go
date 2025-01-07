package di

import (
	"cloud-crm-backend/internal/core/product/application"
	"cloud-crm-backend/internal/core/product/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/product/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideProductDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.ProductRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.ProductUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.ProductHandler, handler)
}
