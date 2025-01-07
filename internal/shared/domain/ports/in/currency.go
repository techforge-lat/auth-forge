package in

import "cloud-crm-backend/internal/core/currency/domain"

type CurrencyUseCase interface {
	UseCaseCommand[domain.CurrencyCreateRequest, domain.CurrencyUpdateRequest]
	UseCaseQuery[domain.Currency]
}
