package out

import "cloud-crm-backend/internal/core/invoicepayment/domain"

type InvoicePaymentRepository interface {
	RepositoryTx[InvoicePaymentRepository]
	RepositoryCommand[domain.InvoicePaymentCreateRequest, domain.InvoicePaymentUpdateRequest]
	RepositoryQuery[domain.InvoicePayment]
}
