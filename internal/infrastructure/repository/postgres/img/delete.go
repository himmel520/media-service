package imgRepo

import (
	"context"
	"errors"

	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ImgRepo) Delete(ctx context.Context, qe repository.Querier, id int) error {
	var pgErr *pgconn.PgError

	cmdTag, err := qe.Exec(ctx, `
	delete from logos 
		where id = $1`, id)
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.FKViolation {
			return repoerr.ErrLogoDependency
		}
	}

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrLogoNotFound
	}

	return err
}
