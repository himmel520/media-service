package tg

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/media-service/pkg/log"
)

func (h *Handler) V1AdminTgsPost(ctx context.Context, req *api.TgPost) (api.V1AdminTgsPostRes, error) {
	tg, err := h.uc.Create(ctx, &entity.Tg{
		Url:   req.URL.String(),
		Title: req.GetTitle(),
	})

	switch {
	case errors.Is(err, repoerr.ErrTGExist):
		return &api.V1AdminTgsPostConflict{Message: err.Error()}, nil
	case err != nil:
		// вот так думаешь использовать норм ? это тоже глабальный инстанс
		// мне так захотелось сделать чтобы избавить хэндлеры от лога в структуре
		// запись можно сократить например до log.Log.Error (это лютая привычка но я бы назвал не принт а console)
		// что бы было console.Log, но это глупо)
		// print.Logger.Error(err)
		log.Info(err)
		log.Error(err)
		// h.log.Error(err)
		return nil, err
	}

	return entity.TgToApi(tg), nil
}
