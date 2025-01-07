package in

import "cloud-crm-backend/internal/core/contractproduct/domain"

type ContractProductUseCase interface {
	UseCaseCommand[domain.ContractProductCreateRequest, domain.ContractProductUpdateRequest]
	UseCaseQuery[domain.ContractProduct]
}
