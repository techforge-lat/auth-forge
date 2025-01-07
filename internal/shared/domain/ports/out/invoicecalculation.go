package out

import "cloud-crm-backend/internal/core/invoicecalculation/domain"

type InvoiceCalculationRepository interface {
	RepositoryTx[InvoiceCalculationRepository]
	RepositoryCommand[domain.InvoiceCalculationCreateRequest, domain.InvoiceCalculationUpdateRequest]
	RepositoryQuery[domain.InvoiceCalculation]
}
