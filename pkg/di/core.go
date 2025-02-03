package di

import (
	"auth-forge/pkg/database"
	"auth-forge/pkg/dependency"

	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/linkit"
)

func ProvideDependencies(container *linkit.DependencyContainer) error {
	db, err := linkit.Resolve[*database.Adapter](container, dependency.Database)
	if err != nil {
		return errortrace.OnError(err)
	}

	provideAppDependencies(container, db)
	provideTenantDependencies(container, db)

	return nil
}
