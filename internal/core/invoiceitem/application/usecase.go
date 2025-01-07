package application

import (
	"cloud-crm-backend/internal/core/invoiceitem/domain"
	"cloud-crm-backend/internal/shared/domain/ports/out"
	"context"

	"github.com/techforge-lat/dafi/v2"
	"github.com/techforge-lat/errortrace/v2"
)

type UseCase struct {
	repo out.InvoiceItemRepository
}

func NewUseCase(repo out.InvoiceItemRepository) UseCase {
	return UseCase{repo: repo}
}

func (uc UseCase) Create(ctx context.Context, entity domain.InvoiceItemCreateRequest) error {
	if err := entity.Validate(); err != nil {
		return errortrace.OnError(err)
	}

	err := uc.repo.Create(ctx, entity)
	if err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (uc UseCase) Update(ctx context.Context, entity domain.InvoiceItemUpdateRequest, filters ...dafi.Filter) error {
	if err := entity.Validate(); err != nil {
		return errortrace.OnError(err)
	}

	err := uc.repo.Update(ctx, entity)
	if err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (uc UseCase) Delete(ctx context.Context, filters ...dafi.Filter) error {
	err := uc.repo.Delete(ctx, filters...)
	if err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (uc UseCase) FindOne(ctx context.Context, criteria dafi.Criteria) (domain.InvoiceItem, error) {
	result, err := uc.repo.FindOne(ctx, criteria)
	if err != nil {
		return domain.InvoiceItem{}, errortrace.OnError(err)
	}

	return result, nil
}

func (uc UseCase) FindAll(ctx context.Context, criteria dafi.Criteria) ([]domain.InvoiceItem, error) {
	result, err := uc.repo.FindAll(ctx, criteria)
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	return result, nil
}
