package in

import "auth-forge/internal/core/app/domain"

type AppUseCase interface {
	UseCaseCommand[domain.AppCreateRequest, domain.AppUpdateRequest]
	UseCaseQuery[domain.App]
}
