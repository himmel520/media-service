package adRepo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
)

func (r *AdRepo) GetByID(ctx context.Context, qe repository.Querier, id int) (*entity.AdvResponse, error) {
	adv := &entity.AdvResponse{}
	err := qe.QueryRow(ctx, `
	SELECT
		adv.id,
		colors.hex AS color,
		logos.url AS logo_url,
		tg.url AS tg_url,
		adv.post,
		adv.title,
		adv.description,
		adv.priority
	FROM adv
	JOIN logos ON adv.logos_id = logos.id
	JOIN colors ON adv.colors_id = colors.id
	JOIN tg ON adv.tg_id = tg.id
		WHERE adv.id = $1;`, id).Scan(
		&adv.ID, &adv.Color, &adv.LogoUrl, &adv.TgUrl, &adv.Post,
		&adv.Title, &adv.Description, &adv.Priority)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrAdvNotFound
	}

	return adv, err
}

func (r *AdRepo) GetAllWithFilter(ctx context.Context, qe repository.Querier, limit, offset int, posts []string, priority []string) ([]*entity.AdvResponse, error) {
	query := `
	SELECT
		adv.id,
		colors.hex AS color,
		logos.url AS logo_url,
		tg.url AS tg_url,
		adv.post,
		adv.title,
		adv.description,
		adv.priority
	FROM adv
	JOIN logos ON adv.logos_id = logos.id
	JOIN colors ON adv.colors_id = colors.id
	JOIN tg ON adv.tg_id = tg.id
	%v
	ORDER BY adv.title ASC
	LIMIT $1 OFFSET $2`

	filter := fmt.Sprintf(`
	WHERE adv.post in ('%v') AND adv.priority in (%v)`,
		strings.Join(posts, "', '"), strings.Join(priority, ", "))

	rows, err := qe.Query(ctx, fmt.Sprintf(query, filter), limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	advs := []*entity.AdvResponse{}
	for rows.Next() {
		adv := &entity.AdvResponse{}

		if err := rows.Scan(&adv.ID, &adv.Color, &adv.LogoUrl, &adv.TgUrl, &adv.Post,
			&adv.Title, &adv.Description, &adv.Priority); err != nil {
			return nil, err
		}

		advs = append(advs, adv)
	}

	if len(advs) == 0 {
		return nil, repoerr.ErrAdvNotFound
	}

	return advs, err
}
