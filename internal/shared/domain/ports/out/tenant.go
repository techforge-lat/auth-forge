package out

import "cloud-crm-backend/internal/core/tenant/domain"

type TenantRepository interface {
	RepositoryTx[TenantRepository]
	RepositoryCommand[domain.TenantCreateRequest, domain.TenantUpdateRequest]
	RepositoryQuery[domain.Tenant]
}
