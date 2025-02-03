package out

import "auth-forge/internal/core/app/domain"

type AppRepository interface {
	RepositoryTx[AppRepository]
	RepositoryCommand[domain.AppCreateRequest, domain.AppUpdateRequest]
	RepositoryQuery[domain.App]
}
