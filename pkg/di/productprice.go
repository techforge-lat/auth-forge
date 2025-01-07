package di

import (
	"cloud-crm-backend/internal/core/productprice/application"
	"cloud-crm-backend/internal/core/productprice/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/productprice/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideProductPriceDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.ProductPriceRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.ProductPriceUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.ProductPriceHandler, handler)
}
