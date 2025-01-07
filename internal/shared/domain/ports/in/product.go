package in

import "cloud-crm-backend/internal/core/product/domain"

type ProductUseCase interface {
	UseCaseCommand[domain.ProductCreateRequest, domain.ProductUpdateRequest]
	UseCaseQuery[domain.Product]
}
