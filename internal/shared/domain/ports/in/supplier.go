package in

import "cloud-crm-backend/internal/core/supplier/domain"

type SupplierUseCase interface {
	UseCaseCommand[domain.SupplierCreateRequest, domain.SupplierUpdateRequest]
	UseCaseQuery[domain.Supplier]
}
