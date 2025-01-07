package in

import "cloud-crm-backend/internal/core/invoicecalculation/domain"

type InvoiceCalculationUseCase interface {
	UseCaseCommand[domain.InvoiceCalculationCreateRequest, domain.InvoiceCalculationUpdateRequest]
	UseCaseQuery[domain.InvoiceCalculation]
}
