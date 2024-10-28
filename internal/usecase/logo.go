package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache/errcache"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository"

	"github.com/sirupsen/logrus"
)

const logoCachePrefix = "logo:*"

type LogoCache interface {
	Get(ctx context.Context, key string) (string, error)
}
type LogoUsecase struct {
	repo  repository.LogoRepo
	cache LogoCache
	log   *logrus.Logger
}

func NewLogoUsecase(repo repository.LogoRepo, cache *cache.Cache, log *logrus.Logger) *LogoUsecase {
	return &LogoUsecase{repo: repo, cache: cache.Client, log: log}
}

func (uc *LogoUsecase) Add(ctx context.Context, logo *entity.Logo) (*entity.LogoResp, error) {
	return uc.repo.Add(ctx, logo)
}

func (uc *LogoUsecase) Update(ctx context.Context, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error) {
	return uc.repo.Update(ctx, id, logo)
}

func (uc *LogoUsecase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *LogoUsecase) GetByID(ctx context.Context, id int) (*entity.LogoResp, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *LogoUsecase) GetAll(ctx context.Context) ([]*entity.LogoResp, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *LogoUsecase) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.LogosResp, error) {
	key := generateCacheKeyLogo(limit, offset)

	val, err := uc.cache.Get(ctx, key)

	if !errors.Is(err, errcache.ErrKeyNotFound) {
		uc.log.Error(err)
	}

	logos := make(map[int]*entity.Logo)
	err = json.Unmarshal([]byte(val), &logos)
	if err != nil {
		logos, err := uc.repo.GetAllWithPagination(ctx, limit, offset)
		if err != nil {
			return nil, err
		}

		return &entity.LogosResp{
			Logos: logos,
			Total: len(logos),
		}, err

	}

	return &entity.LogosResp{
		Logos: logos,
		Total: len(logos),
	}, err

}

func generateCacheKeyLogo(limit, offset int) string {
	key := fmt.Sprintf("%d:%d:%s:%s", limit, offset)

	// Создаем хеш
	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return advCachePrefix + hash
}
