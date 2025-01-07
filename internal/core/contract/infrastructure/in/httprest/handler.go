package httprest

import (
	"cloud-crm-backend/internal/core/contract/domain"
	"cloud-crm-backend/internal/shared/domain/ports/in"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/techforge-lat/dafi/v2"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/errortrace/v2/errtype"
	"github.com/techforge-lat/rapi"
)

type Handler struct {
	useCase in.ContractUseCase
}

func New(useCase in.ContractUseCase) Handler {
	return Handler{useCase: useCase}
}

func (h Handler) Create(c echo.Context) error {
	entity := domain.ContractCreateRequest{}

	if err := c.Bind(&entity); err != nil {
		return errortrace.OnError(err).WithCode(errtype.UnprocessableEntity)
	}

	if err := h.useCase.Create(c.Request().Context(), entity); err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusCreated, rapi.Created(entity))
}

func (h Handler) UpdateByID(c echo.Context) error {
	entity := domain.ContractUpdateRequest{}

	if err := c.Bind(&entity); err != nil {
		return errortrace.OnError(err).WithCode(errtype.UnprocessableEntity)
	}

	if err := h.useCase.Update(c.Request().Context(), entity, dafi.FilterBy("id", dafi.Equal, c.Param("id"))...); err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Updated())
}

func (h Handler) Update(c echo.Context) error {
	entity := domain.ContractUpdateRequest{}

	if err := c.Bind(&entity); err != nil {
		return errortrace.OnError(err).WithCode(errtype.UnprocessableEntity)
	}

	criteria, err := dafi.NewQueryParser().Parse(c.QueryParams())
	if err != nil {
		return errortrace.OnError(err)
	}

	if err := h.useCase.Update(c.Request().Context(), entity, criteria.Filters...); err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Updated())
}

func (h Handler) DeleteByID(c echo.Context) error {
	if err := h.useCase.Delete(c.Request().Context(), dafi.FilterBy("id", dafi.Equal, c.Param("id"))...); err != nil {
		return errortrace.OnError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h Handler) Delete(c echo.Context) error {
	criteria, err := dafi.NewQueryParser().Parse(c.QueryParams())
	if err != nil {
		return errortrace.OnError(err)
	}

	if err := h.useCase.Delete(c.Request().Context(), criteria.Filters...); err != nil {
		return errortrace.OnError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h Handler) FindOneByID(c echo.Context) error {
	result, err := h.useCase.FindOne(c.Request().Context(), dafi.Where("id", dafi.Equal, c.Param("id")))
	if err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Ok(result))
}

func (h Handler) FindOne(c echo.Context) error {
	criteria, err := dafi.NewQueryParser().Parse(c.QueryParams())
	if err != nil {
		return errortrace.OnError(err)
	}

	result, err := h.useCase.FindOne(c.Request().Context(), criteria)
	if err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Ok(result))
}

func (h Handler) FindAll(c echo.Context) error {
	criteria, err := dafi.NewQueryParser().Parse(c.QueryParams())
	if err != nil {
		return errortrace.OnError(err)
	}

	result, err := h.useCase.FindAll(c.Request().Context(), criteria)
	if err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Ok(result))
}
