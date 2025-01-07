package out

import "cloud-crm-backend/internal/core/contract/domain"

type ContractRepository interface {
	RepositoryTx[ContractRepository]
	RepositoryCommand[domain.ContractCreateRequest, domain.ContractUpdateRequest]
	RepositoryQuery[domain.Contract]
}
