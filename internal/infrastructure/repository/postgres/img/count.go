package imgRepo

import (
	"context"

	"github.com/himmel520/media-service/internal/infrastructure/repository"
)

func (r *ImgRepo) Count(ctx context.Context, qe repository.Querier) (int, error) {
	var count int
	err := qe.QueryRow(ctx, `SELECT COUNT(*) FROM images`).Scan(&count)
	return count, err
}
