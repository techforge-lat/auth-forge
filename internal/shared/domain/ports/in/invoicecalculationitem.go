package in

import "cloud-crm-backend/internal/core/invoicecalculationitem/domain"

type InvoiceCalculationItemUseCase interface {
	UseCaseCommand[domain.InvoiceCalculationItemCreateRequest, domain.InvoiceCalculationItemUpdateRequest]
	UseCaseQuery[domain.InvoiceCalculationItem]
}
