package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// TODO: убрать зависимость интерфейса от pgx

type (
	DBTX struct {
		qe Querier
	}

	Querier interface {
		Begin(ctx context.Context) (pgx.Tx, error)
		Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	}

	QuerierTX interface {
		Querier
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}

	TransactionFunc func(ctx context.Context, qe Querier) error
)

func NewDBTX(db Querier) *DBTX {
	return &DBTX{qe: db}
}

func (d *DBTX) DB() Querier {
	return d.qe
}

func (d *DBTX) InTransaction(ctx context.Context, fn TransactionFunc) error {
	tx, err := d.qe.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	err = fn(ctx, tx)
	if err != nil {
		// TODO: убоать pgx.ErrTxClosed
		if rbErr := tx.Rollback(ctx); rbErr != nil && !errors.Is(err, pgx.ErrTxClosed) {
			return fmt.Errorf("rollback err: %w; err: %w", rbErr, err)
		}
		return err
	}

	return tx.Commit(ctx)
}
