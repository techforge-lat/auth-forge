package out

import "cloud-crm-backend/internal/core/paymentaccount/domain"

type PaymentAccountRepository interface {
	RepositoryTx[PaymentAccountRepository]
	RepositoryCommand[domain.PaymentAccountCreateRequest, domain.PaymentAccountUpdateRequest]
	RepositoryQuery[domain.PaymentAccount]
}
