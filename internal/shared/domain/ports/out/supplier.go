package out

import "cloud-crm-backend/internal/core/supplier/domain"

type SupplierRepository interface {
	RepositoryTx[SupplierRepository]
	RepositoryCommand[domain.SupplierCreateRequest, domain.SupplierUpdateRequest]
	RepositoryQuery[domain.Supplier]
}
