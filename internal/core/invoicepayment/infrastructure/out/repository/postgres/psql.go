package postgres

import (
	"cloud-crm-backend/internal/core/invoicepayment/domain"
	"cloud-crm-backend/internal/shared/domain/ports/out"
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/techforge-lat/dafi/v2"
	"github.com/techforge-lat/errortrace/v2"
)

type Repository struct {
	db out.Database
	tx out.Tx
}

func NewRepository(db out.Database) Repository {
	return Repository{db: db}
}

// WithTx returns a new instance of the repository with the transaction set
func (r Repository) WithTx(tx out.Transaction) out.InvoicePaymentRepository {
	return Repository{
		db: r.db,
		tx: tx.GetTx(),
	}
}

func (r Repository) Create(ctx context.Context, entity domain.InvoicePaymentCreateRequest) error {
	result, err := insertQuery.WithValues(entity.InvoiceID, entity.PaymentAccountID, entity.PaymentMethodID, entity.Amount, entity.Notes, entity.ExchangeRate, entity.PaymentDate, entity.Hash, entity.ReferenceCode, entity.IsDetraction, entity.PenUsdExchangeRate, entity.CreatedAt).ToSQL()
	if err != nil {
		return errortrace.OnError(err)
	}

	if _, err := r.conn().Exec(ctx, result.Sql, result.Args...); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (r Repository) Update(ctx context.Context, entity domain.InvoicePaymentUpdateRequest, filters ...dafi.Filter) error {
	if !entity.UpdatedAt.Valid {
		entity.UpdatedAt.SetValid(time.Now())
	}

	result, err := updateQuery.WithValues(entity.InvoiceID, entity.PaymentAccountID, entity.PaymentMethodID, entity.Amount, entity.Notes, entity.ExchangeRate, entity.PaymentDate, entity.Hash, entity.ReferenceCode, entity.IsDetraction, entity.PenUsdExchangeRate, entity.UpdatedAt).Where(filters...).ToSQL()
	if err != nil {
		return errortrace.OnError(err)
	}

	if _, err := r.conn().Exec(ctx, result.Sql, result.Args...); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, filters ...dafi.Filter) error {
	result, err := deleteQuery.Where(filters...).ToSQL()
	if err != nil {
		return errortrace.OnError(err)
	}

	if _, err := r.conn().Exec(ctx, result.Sql, result.Args...); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (r Repository) FindOne(ctx context.Context, criteria dafi.Criteria) (domain.InvoicePayment, error) {
	result, err := selectQuery.Where(criteria.Filters...).OrderBy(criteria.Sorts...).Limit(1).RequiredColumns(criteria.SelectColumns...).ToSQL()
	if err != nil {
		return domain.InvoicePayment{}, errortrace.OnError(err)
	}

	var m domain.InvoicePayment
	if err := pgxscan.Get(ctx, r.conn(), &m, result.Sql, result.Args...); err != nil {
		return domain.InvoicePayment{}, errortrace.OnError(err)
	}

	return m, nil
}

func (r Repository) FindAll(ctx context.Context, criteria dafi.Criteria) ([]domain.InvoicePayment, error) {
	result, err := selectQuery.Where(criteria.Filters...).OrderBy(criteria.Sorts...).Limit(criteria.Pagination.PageSize).Page(criteria.Pagination.PageNumber).RequiredColumns(criteria.SelectColumns...).ToSQL()
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	var ms []domain.InvoicePayment
	if err := pgxscan.Select(ctx, r.conn(), &ms, result.Sql, result.Args...); err != nil {
		return nil, errortrace.OnError(err)
	}

	return ms, nil
}

// conn returns the database connection to use
// if there is a transaction, it returns the transaction connection
func (r Repository) conn() out.Database {
	if r.tx != nil {
		return r.tx
	}

	return r.db
}
