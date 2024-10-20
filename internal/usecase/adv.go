package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache/errcache"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository"

	"github.com/sirupsen/logrus"
)

type AdvUsecase struct {
	repo  repository.AdvRepo
	cache cache.AdvCache
	log   *logrus.Logger
}

func NewAdvUsecase(repo repository.AdvRepo, cache cache.AdvCache, log *logrus.Logger) *AdvUsecase {
	return &AdvUsecase{repo: repo, cache: cache, log: log}
}

func (s *AdvUsecase) Add(ctx context.Context, adv *entity.Adv) (*entity.AdvResponse, error) {
	id, err := s.repo.Add(ctx, adv)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, id)
}

func (s *AdvUsecase) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *AdvUsecase) Update(ctx context.Context, id int, adv *entity.AdvUpdate) (*entity.AdvResponse, error) {
	if err := s.repo.Update(ctx, id, adv); err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, id)
}

func (s *AdvUsecase) GetAllWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*entity.AdvResponse, error) {
	key := s.generateCacheKey(limit, offset, posts, priority)

	advs, err := s.cache.Get(ctx, key)
	if err != nil {
		if !errors.Is(err, errcache.ErrKeyNotFound) {
			s.log.Error(err)
		}

		// идем в бд
		advs, err = s.repo.GetAllWithFilter(ctx, limit, offset, posts, priority)
		if err != nil {
			return nil, err
		}

		// сохраняем в кэш
		if err := s.cache.Set(ctx, key, advs); err != nil {
			s.log.Error(err)
		}
	}

	return advs, nil
}

func (s *AdvUsecase) DeleteCache(ctx context.Context) error {
	return s.cache.Delete(ctx)
}

func (s *AdvUsecase) generateCacheKey(limit, offset int, posts, priority []string) string {
	key := fmt.Sprintf("%d:%d:%s:%s", limit, offset, strings.Join(posts, ","), strings.Join(priority, ","))

	// Создаем хеш
	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return "advs:" + hash
}
