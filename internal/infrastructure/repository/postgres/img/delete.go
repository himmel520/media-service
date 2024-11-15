package imgRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ImgRepo) Delete(ctx context.Context, qe repository.Querier, id int) error {
	query, args, err := squirrel.Delete("images").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	var pgErr *pgconn.PgError
	cmdTag, err := qe.Exec(ctx, query, args...)
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.FKViolation {
			return repoerr.ErrImageDependencyExist
		}
	}

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrImageNotFound
	}

	return err
}
