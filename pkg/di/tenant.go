package di

import (
	"auth-forge/internal/core/tenant/application"
	"auth-forge/internal/core/tenant/infrastructure/in/httprest"
	"auth-forge/internal/core/tenant/infrastructure/out/repository/postgres"
	"auth-forge/pkg/database"
	"auth-forge/pkg/dependency"

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
