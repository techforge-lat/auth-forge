package out

import "auth-forge/internal/core/tenant/domain"

type TenantRepository interface {
	RepositoryTx[TenantRepository]
	RepositoryCommand[domain.TenantCreateRequest, domain.TenantUpdateRequest]
	RepositoryQuery[domain.Tenant]
}
