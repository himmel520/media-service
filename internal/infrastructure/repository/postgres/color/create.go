package colorRepo

import (
	"context"
	"errors"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ColorRepo) Add(ctx context.Context, qe repository.Querier, color *entity.Color) (*entity.ColorResp, error) {
	newColor := &entity.ColorResp{}

	err := qe.QueryRow(ctx, `
	insert into colors 
		(title, hex) 
	values 
		($1, $2) 
	returning *`, color.Title, color.Hex).Scan(&newColor.ID, &newColor.Title, &newColor.Hex)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrColorHexExist
	}

	return newColor, err
}
