package out

import "cloud-crm-backend/internal/core/invoicecalculationitem/domain"

type InvoiceCalculationItemRepository interface {
	RepositoryTx[InvoiceCalculationItemRepository]
	RepositoryCommand[domain.InvoiceCalculationItemCreateRequest, domain.InvoiceCalculationItemUpdateRequest]
	RepositoryQuery[domain.InvoiceCalculationItem]
}
