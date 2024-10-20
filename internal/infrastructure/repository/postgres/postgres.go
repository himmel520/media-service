package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPG(pgConn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), pgConn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
