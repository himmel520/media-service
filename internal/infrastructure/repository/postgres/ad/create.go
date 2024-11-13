package adRepo

import (
	"context"
	"errors"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AdRepo) Add(ctx context.Context, qe repository.Querier, adv *entity.Adv) (int, error) {
	var id int
	err := qe.QueryRow(ctx, `
	insert into adv
		(logos_id, colors_id, tg_id, post, title, description, priority)
	values
		($1, $2, $3, $4, $5, $6, $7)
	returning id;`,
		adv.LogoID, adv.ColorID, adv.TgID, adv.Post, adv.Title, adv.Description, adv.Priority).Scan(&id)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.FKViolation {
			return 0, repoerr.ErrAdvDependencyNotExist
		}
	}

	return id, err
}
