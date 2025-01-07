package di

import (
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/linkit"
)

func ProvideDependencies(container *linkit.DependencyContainer) error {
	db, err := linkit.Resolve[*database.Adapter](container, dependency.Database)
	if err != nil {
		return errortrace.OnError(err)
	}

	provideTenantDependencies(container, db)
	provideSupplierDependencies(container, db)

	return nil
}
