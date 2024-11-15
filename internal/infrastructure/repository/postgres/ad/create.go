package adRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AdRepo) Create(ctx context.Context, qe repository.Querier, adv *entity.Adv) (*entity.AdvResp, error) {
	query, args, err := squirrel.Insert("adv").
		Columns(
			"images_id",
			"colors_id",
			"tg_id",
			"post",
			"title",
			"description",
			"priority").
		Values(
			adv.ImageID,
			adv.ColorID,
			adv.TgID,
			adv.Post,
			adv.Title,
			adv.Description,
			adv.Priority).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(`returning id`).
		ToSql()
	if err != nil {
		return nil, err
	}

	var id int
	err = qe.QueryRow(ctx, query, args...).Scan(&id)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgErr) {
			if pgErr.Code == repoerr.FKViolation {
				return nil, repoerr.ErrAdvDependencyNotExist
			}
		}
		return nil, err
	}

	return r.GetByID(ctx, qe, id)
}
