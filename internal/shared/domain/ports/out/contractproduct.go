package out

import "cloud-crm-backend/internal/core/contractproduct/domain"

type ContractProductRepository interface {
	RepositoryTx[ContractProductRepository]
	RepositoryCommand[domain.ContractProductCreateRequest, domain.ContractProductUpdateRequest]
	RepositoryQuery[domain.ContractProduct]
}
