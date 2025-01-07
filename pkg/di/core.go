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
	provideProductDependencies(container, db)
	provideProductPriceDependencies(container, db)
	provideContractDependencies(container, db)
	provideContractProductDependencies(container, db)
	provideInvoiceDependencies(container, db)
	provideInvoiceItemDependencies(container, db)
	provideInvoicePaymentDependencies(container, db)
	provideInvoiceCalculationDependencies(container, db)
	provideInvoiceCalculationItemDependencies(container, db)
	provideCurrencyDependencies(container, db)
	providePaymentAccountDependencies(container, db)

	return nil
}
