package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/uoffer/mediaAd/internal/cache"
	"github.com/himmel520/uoffer/mediaAd/internal/models"

	"github.com/sirupsen/logrus"
)

//go:generate mockery --all

type AdvRepo interface {
	Add(ctx context.Context, adv *models.Adv) (int, error)
	GetByID(ctx context.Context, id int) (*models.AdvResponse, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, adv *models.AdvUpdate) error
	GetAllWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*models.AdvResponse, error)
}

type AdvCache interface {
	Set(ctx context.Context, key string, advs []*models.AdvResponse) error
	Get(ctx context.Context, key string) ([]*models.AdvResponse, error)
	Delete(ctx context.Context) error
}

type AdvService struct {
	repo  AdvRepo
	cache AdvCache
	log   *logrus.Logger
}

func NewAdvService(repo AdvRepo, cache AdvCache, log *logrus.Logger) *AdvService {
	return &AdvService{repo: repo, cache: cache, log: log}
}

func (s *AdvService) Add(ctx context.Context, adv *models.Adv) (*models.AdvResponse, error) {
	id, err := s.repo.Add(ctx, adv)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, id)
}

func (s *AdvService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *AdvService) Update(ctx context.Context, id int, adv *models.AdvUpdate) (*models.AdvResponse, error) {
	if err := s.repo.Update(ctx, id, adv); err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, id)
}

func (s *AdvService) GetAllWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*models.AdvResponse, error) {
	key := s.generateCacheKey(limit, offset, posts, priority)

	advs, err := s.cache.Get(ctx, key)
	if err != nil {
		if !errors.Is(err, cache.ErrKeyNotFound) {
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

func (s *AdvService) DeleteCache(ctx context.Context) error {
	return s.cache.Delete(ctx)
}

func (s *AdvService) generateCacheKey(limit, offset int, posts, priority []string) string {
	key := fmt.Sprintf("%d:%d:%s:%s", limit, offset, strings.Join(posts, ","), strings.Join(priority, ","))

	// Создаем хеш
	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return "advs:" + hash
}
