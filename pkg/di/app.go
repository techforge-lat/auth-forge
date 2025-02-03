package di

import (
	"auth-forge/internal/core/app/application"
	"auth-forge/internal/core/app/infrastructure/in/httprest"
	"auth-forge/internal/core/app/infrastructure/out/repository/postgres"
	"auth-forge/pkg/database"
	"auth-forge/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func provideAppDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.AppRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.AppUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.AppHandler, handler)
}
