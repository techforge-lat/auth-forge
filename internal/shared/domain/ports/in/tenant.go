package in

import "cloud-crm-backend/internal/core/tenant/domain"

type TenantUseCase interface {
	UseCaseCommand[domain.TenantCreateRequest, domain.TenantUpdateRequest]
	UseCaseQuery[domain.Tenant]
}
