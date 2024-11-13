package adRepo

import (
	"context"

	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (r *AdRepo) Delete(ctx context.Context, qe repository.Querier, id int) error {
	cmdTag, err := qe.Exec(ctx, `delete from adv where id = $1`, id)
	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrAdvNotFound
	}
	return err

}
