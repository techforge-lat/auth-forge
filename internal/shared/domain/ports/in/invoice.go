package in

import "cloud-crm-backend/internal/core/invoice/domain"

type InvoiceUseCase interface {
	UseCaseCommand[domain.InvoiceCreateRequest, domain.InvoiceUpdateRequest]
	UseCaseQuery[domain.Invoice]
}
