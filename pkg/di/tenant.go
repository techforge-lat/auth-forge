package di

import (
	"cloud-crm-backend/internal/core/tenant/application"
	"cloud-crm-backend/internal/core/tenant/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/tenant/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideTenantDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.TenantRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.TenantUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.TenantHandler, handler)
}
