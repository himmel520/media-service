package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache/errcache"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository"

	"github.com/sirupsen/logrus"
)

type AdvCache interface {
	Set(ctx context.Context, key string, advs any) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context) error
}

type AdvUsecase struct {
	repo  repository.AdvRepo
	cache AdvCache
	log   *logrus.Logger
}

func NewAdvUsecase(repo repository.AdvRepo, cache *cache.Cache, log *logrus.Logger) *AdvUsecase {
	return &AdvUsecase{repo: repo, cache: cache.Client, log: log}
}

func (uc *AdvUsecase) Add(ctx context.Context, adv *entity.Adv) (*entity.AdvResponse, error) {
	id, err := uc.repo.Add(ctx, adv)
	if err != nil {
		return nil, err
	}

	return uc.repo.GetByID(ctx, id)
}

func (uc *AdvUsecase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *AdvUsecase) Update(ctx context.Context, id int, adv *entity.AdvUpdate) (*entity.AdvResponse, error) {
	if err := uc.repo.Update(ctx, id, adv); err != nil {
		return nil, err
	}

	return uc.repo.GetByID(ctx, id)
}

func (uc *AdvUsecase) GetAllWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*entity.AdvResponse, error) {
	key := uc.generateCacheKey(limit, offset, posts, priority)

	val, err := uc.cache.Get(ctx, key)

	// хочется пояснить но потом надо убрать, игнорируются все ошибки из кэша так как не имеет смысла прерывать запрос
	// таким образом увеличиться надежность
	if !errors.Is(err, errcache.ErrKeyNotFound) {
		uc.log.Error(err)
	}

	advs := []*entity.AdvResponse{}
	err = json.Unmarshal([]byte(val), &advs)

	if err != nil {
		advs, err = uc.repo.GetAllWithFilter(ctx, limit, offset, posts, priority)
		if err != nil {
			return nil, err
		}

		if err := uc.cache.Set(ctx, key, advs); err != nil {
			uc.log.Error(err)
		}
	}
	// TODO: надо протестить руками
	// TODO: решается дженериками для методов, но так как они не работают из Get возвращается any
	return advs, nil
}

func (uc *AdvUsecase) DeleteCache(ctx context.Context) error {
	return uc.cache.Delete(ctx)
}

func (uc *AdvUsecase) generateCacheKey(limit, offset int, posts, priority []string) string {
	key := fmt.Sprintf("%d:%d:%s:%s", limit, offset, strings.Join(posts, ","), strings.Join(priority, ","))

	// Создаем хеш
	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return "advs:" + hash
}
