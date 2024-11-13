package tgRepo

import (
	"context"
	"errors"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *TgRepo) Add(ctx context.Context, qe repository.Querier, tg *entity.TG) (*entity.TGResp, error) {
	newTG := &entity.TGResp{}

	err := qe.QueryRow(ctx, `
	insert into tg 
		(url, title) 
	values 
		($1, $2) 
	returning *`, tg.Url, tg.Title).Scan(&newTG.ID, &newTG.Title, &newTG.Url)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrTGExist
	}

	return newTG, err
}
