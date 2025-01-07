package out

import "cloud-crm-backend/internal/core/currency/domain"

type CurrencyRepository interface {
	RepositoryTx[CurrencyRepository]
	RepositoryCommand[domain.CurrencyCreateRequest, domain.CurrencyUpdateRequest]
	RepositoryQuery[domain.Currency]
}
