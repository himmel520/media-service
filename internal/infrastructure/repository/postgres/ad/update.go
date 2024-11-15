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

func (r *AdRepo) Update(ctx context.Context, qe repository.Querier, id int, adv *entity.AdvUpdate) (*entity.AdvResp, error) {
	builder := squirrel.Update("adv").
		Where(squirrel.Eq{"id": id}).
		Suffix(`returning id`).
		PlaceholderFormat(squirrel.Dollar)

	if adv.ImageID.Set {
		builder = builder.Set("images_id", adv.ImageID.Value)
	}

	if adv.ColorID.Set {
		builder = builder.Set("colors_id", adv.ColorID.Value)
	}

	if adv.TgID.Set {
		builder = builder.Set("tg_id", adv.TgID.Value)
	}

	if adv.Post.Set {
		builder = builder.Set("post", adv.Post.Value)
	}

	if adv.Title.Set {
		builder = builder.Set("title", adv.Title.Value)
	}

	if adv.Description.Set {
		builder = builder.Set("description", adv.Description.Value)
	}

	if adv.Priority.Set {
		builder = builder.Set("post", adv.Priority.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	cmdTag, err := qe.Exec(ctx, query, args...)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == repoerr.FKViolation {
			return nil, repoerr.ErrAdvDependencyNotExist
		}
		return nil, err
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, repoerr.ErrAdvNotFound
	}

	return r.GetByID(ctx, qe, id)
}
