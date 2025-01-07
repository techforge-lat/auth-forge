package out

import "cloud-crm-backend/internal/core/product/domain"

type ProductRepository interface {
	RepositoryTx[ProductRepository]
	RepositoryCommand[domain.ProductCreateRequest, domain.ProductUpdateRequest]
	RepositoryQuery[domain.Product]
}
