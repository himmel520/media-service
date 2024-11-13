package colorRepo

import (
	"context"

	"github.com/himmel520/media-service/internal/infrastructure/repository"
)

func (r *ColorRepo) Count(ctx context.Context, qe repository.Querier) (int, error) {
	var count int
	err := qe.QueryRow(ctx, `SELECT COUNT(*) FROM colors;`).Scan(&count)
	return count, err
}
