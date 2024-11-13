package imgRepo

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

func (r *ImgRepo) Update(ctx context.Context, qe repository.Querier, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error) {
	var keys []string
	var values []interface{}
	if logo.Url != nil {
		keys = append(keys, "url=$1")
		values = append(values, logo.Url)
	}

	if logo.Title != nil {
		keys = append(keys, fmt.Sprintf("title=$%d", len(values)+1))
		values = append(values, logo.Title)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update logos 
		set %v 
	where id = $%v
	returning *;`, strings.Join(keys, ", "), len(values))

	newLogo := &entity.LogoResp{}
	err := qe.QueryRow(ctx, query, values...).Scan(&newLogo.ID, &newLogo.Url, &newLogo.Title)

	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, repoerr.ErrLogoNotFound
	case errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint:
		return nil, repoerr.ErrLogoExist
	}

	return newLogo, err
}
