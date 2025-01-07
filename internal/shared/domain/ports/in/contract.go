package in

import "cloud-crm-backend/internal/core/contract/domain"

type ContractUseCase interface {
	UseCaseCommand[domain.ContractCreateRequest, domain.ContractUpdateRequest]
	UseCaseQuery[domain.Contract]
}
