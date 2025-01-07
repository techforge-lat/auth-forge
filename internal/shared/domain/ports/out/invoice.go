package out

import "cloud-crm-backend/internal/core/invoice/domain"

type InvoiceRepository interface {
	RepositoryTx[InvoiceRepository]
	RepositoryCommand[domain.InvoiceCreateRequest, domain.InvoiceUpdateRequest]
	RepositoryQuery[domain.Invoice]
}
