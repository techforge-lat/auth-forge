package in

import "cloud-crm-backend/internal/core/invoicepayment/domain"

type InvoicePaymentUseCase interface {
	UseCaseCommand[domain.InvoicePaymentCreateRequest, domain.InvoicePaymentUpdateRequest]
	UseCaseQuery[domain.InvoicePayment]
}
