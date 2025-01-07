package out

import "cloud-crm-backend/internal/core/invoiceitem/domain"

type InvoiceItemRepository interface {
	RepositoryTx[InvoiceItemRepository]
	RepositoryCommand[domain.InvoiceItemCreateRequest, domain.InvoiceItemUpdateRequest]
	RepositoryQuery[domain.InvoiceItem]
}
