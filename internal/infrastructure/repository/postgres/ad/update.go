package adRepo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AdRepo) Update(ctx context.Context, qe repository.Querier, id int, adv *entity.AdvUpdate) error {
	var keys []string
	var values []interface{}

	if adv.LogoID != nil {
		keys = append(keys, "logos_id=$1")
		values = append(values, adv.LogoID)
	}

	if adv.ColorID != nil {
		keys = append(keys, fmt.Sprintf("colors_id=$%d", len(keys)+1))
		values = append(values, adv.ColorID)
	}

	if adv.TgID != nil {
		keys = append(keys, fmt.Sprintf("tg_id=$%d", len(keys)+1))
		values = append(values, adv.TgID)
	}

	if adv.Post != nil {
		keys = append(keys, fmt.Sprintf("post=$%d", len(keys)+1))
		values = append(values, adv.Post)
	}

	if adv.Title != nil {
		keys = append(keys, fmt.Sprintf("title=$%d", len(keys)+1))
		values = append(values, adv.Title)
	}

	if adv.Description != nil {
		keys = append(keys, fmt.Sprintf("description=$%d", len(keys)+1))
		values = append(values, adv.Description)
	}

	if adv.Priority != nil {
		keys = append(keys, fmt.Sprintf("priority=$%d", len(keys)+1))
		values = append(values, adv.Priority)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update adv 
		set %s 
	where id = $%d`, strings.Join(keys, ", "), len(values))

	var pgErr *pgconn.PgError
	cmdTag, err := qe.Exec(ctx, query, values...)
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.FKViolation {
			return repoerr.ErrAdvDependencyNotExist
		}
	}

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrAdvNotFound
	}

	return err
}
