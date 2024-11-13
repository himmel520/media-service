package colorRepo

import (
	"context"
	"errors"

	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ColorRepo) Delete(ctx context.Context, qe repository.Querier, id int) error {
	cmdTag, err := qe.Exec(ctx, `
	delete from Colors 
		where id = $1`, id)

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrColorNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.FKViolation {
		return repoerr.ErrColorDependencyExist
	}

	return err
}
