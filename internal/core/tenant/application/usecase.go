package application

import (
	"auth-forge/internal/core/tenant/domain"
	"auth-forge/internal/shared/application/code"
	"auth-forge/internal/shared/domain/ports/out"
	"context"

	"github.com/techforge-lat/dafi/v2"
	"github.com/techforge-lat/errortrace/v2"
)

type UseCase struct {
	repo out.TenantRepository
}

func NewUseCase(repo out.TenantRepository) UseCase {
	return UseCase{repo: repo}
}

func (uc UseCase) Create(ctx context.Context, entity domain.TenantCreateRequest) error {
	if err := entity.Validate(); err != nil {
		return errortrace.OnError(err)
	}

	code, err := code.Generate(entity.Name, 5)
	if err != nil {
		return errortrace.OnError(err)
	}
	entity.Code = code

	err = uc.repo.Create(ctx, entity)
	if err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (uc UseCase) Update(ctx context.Context, entity domain.TenantUpdateRequest, filters ...dafi.Filter) error {
	if err := entity.Validate(); err != nil {
		return errortrace.OnError(err)
	}

	err := uc.repo.Update(ctx, entity, filters...)
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

func (uc UseCase) FindOne(ctx context.Context, criteria dafi.Criteria) (domain.Tenant, error) {
	result, err := uc.repo.FindOne(ctx, criteria)
	if err != nil {
		return domain.Tenant{}, errortrace.OnError(err)
	}

	return result, nil
}

func (uc UseCase) FindAll(ctx context.Context, criteria dafi.Criteria) ([]domain.Tenant, error) {
	result, err := uc.repo.FindAll(ctx, criteria)
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	return result, nil
}
