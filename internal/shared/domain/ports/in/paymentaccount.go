package in

import "cloud-crm-backend/internal/core/paymentaccount/domain"

type PaymentAccountUseCase interface {
	UseCaseCommand[domain.PaymentAccountCreateRequest, domain.PaymentAccountUpdateRequest]
	UseCaseQuery[domain.PaymentAccount]
}
