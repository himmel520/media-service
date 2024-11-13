package colorRepo

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

func (r *ColorRepo) Update(ctx context.Context, qe repository.Querier, id int, color *entity.ColorUpdate) (*entity.ColorResp, error) {
	var keys []string
	var values []interface{}
	if color.Title != nil {
		keys = append(keys, "title=$1")
		values = append(values, color.Title)
	}

	if color.Hex != nil {
		keys = append(keys, fmt.Sprintf("hex=$%d", len(values)+1))
		values = append(values, color.Hex)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update colors 
		set %v 
	where id = $%v
	returning *;`, strings.Join(keys, ", "), len(values))

	newColor := &entity.ColorResp{}
	err := qe.QueryRow(ctx, query, values...).Scan(&newColor.ID, &newColor.Title, &newColor.Hex)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrColorNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrColorHexExist
	}

	return newColor, err
}
