package in

import "auth-forge/internal/core/tenant/domain"

type TenantUseCase interface {
	UseCaseCommand[domain.TenantCreateRequest, domain.TenantUpdateRequest]
	UseCaseQuery[domain.Tenant]
}
