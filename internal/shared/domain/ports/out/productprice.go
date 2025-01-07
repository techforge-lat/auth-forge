package out

import "cloud-crm-backend/internal/core/productprice/domain"

type ProductPriceRepository interface {
	RepositoryTx[ProductPriceRepository]
	RepositoryCommand[domain.ProductPriceCreateRequest, domain.ProductPriceUpdateRequest]
	RepositoryQuery[domain.ProductPrice]
}
