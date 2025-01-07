package di

import (
	"cloud-crm-backend/internal/core/paymentaccount/application"
	"cloud-crm-backend/internal/core/paymentaccount/infrastructure/in/httprest"
	"cloud-crm-backend/internal/core/paymentaccount/infrastructure/out/repository/postgres"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"

	"github.com/techforge-lat/linkit"
)

func providePaymentAccountDependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.PaymentAccountRepository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.PaymentAccountUseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.PaymentAccountHandler, handler)
}
