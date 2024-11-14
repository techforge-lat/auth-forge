package database

import (
	"cloud-crm-backend/internal/shared/domain"
	"context"
	"fmt"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/techforge-lat/errortrace/v2"
)

type Adapter struct {
	*pgxpool.Pool
}

func New(conf domain.Configuration) (*Adapter, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		conf.Database.Driver,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
		conf.Database.SSLMode,
	))
	if err != nil {
		return nil, errortrace.OnError(fmt.Errorf("unable to parse config connection: %w", err))
	}

	config.ConnConfig.Tracer = otelpgx.NewTracer(otelpgx.WithIncludeQueryParameters())

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, errortrace.OnError(fmt.Errorf("unable to create connection pool: %w", err))
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		return nil, errortrace.OnError(fmt.Errorf("unable to ping database: %w", err))
	}

	return &Adapter{dbPool}, nil
}
