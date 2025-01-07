package in

import "cloud-crm-backend/internal/core/productprice/domain"

type ProductPriceUseCase interface {
	UseCaseCommand[domain.ProductPriceCreateRequest, domain.ProductPriceUpdateRequest]
	UseCaseQuery[domain.ProductPrice]
}
