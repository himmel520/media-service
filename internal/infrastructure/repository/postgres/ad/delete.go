package adRepo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (r *AdRepo) Delete(ctx context.Context, qe repository.Querier, id int) error {
	query, args, err := squirrel.Delete("adv").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	cmdTag, err := qe.Exec(ctx, query, args...)
	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrAdvNotFound
	}
	return err

}
