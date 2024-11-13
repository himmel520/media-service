package tgRepo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *TgRepo) Update(ctx context.Context, qe repository.Querier, id int, tg *entity.TGUpdate) (*entity.TGResp, error) {
	var keys []string
	var values []interface{}
	if tg.Url != nil {
		keys = append(keys, "url=$1")
		values = append(values, tg.Url)
	}

	if tg.Title != nil {
		keys = append(keys, fmt.Sprintf("title=$%d", len(values)+1))
		values = append(values, tg.Title)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update tg 
		set %v 
	where id = $%v
	returning *;`, strings.Join(keys, ", "), len(values))

	newTG := &entity.TGResp{}
	err := qe.QueryRow(ctx, query, values...).Scan(&newTG.ID, &newTG.Title, &newTG.Url)

	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, repoerr.ErrTGNotFound
	case errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint:
		return nil, repoerr.ErrTGExist
	}

	return newTG, err
}
