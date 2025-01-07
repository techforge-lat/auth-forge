package in

import "cloud-crm-backend/internal/core/invoiceitem/domain"

type InvoiceItemUseCase interface {
	UseCaseCommand[domain.InvoiceItemCreateRequest, domain.InvoiceItemUpdateRequest]
	UseCaseQuery[domain.InvoiceItem]
}
