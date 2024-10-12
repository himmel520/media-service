package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/uoffer/mediaAd/internal/repository"
	"github.com/himmel520/uoffer/mediaAd/models"
)

func (s *Service) AddAdv(ctx context.Context, adv *models.Adv) (*models.AdvResponse, error) {
	id, err := s.repo.AddAdv(ctx, adv)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAdvByID(ctx, id)
}

func (s *Service) DeleteAdv(ctx context.Context, id int) error {
	return s.repo.DeleteAdv(ctx, id)
}

func (s *Service) UpdateAdv(ctx context.Context, id int, adv *models.AdvUpdate) (*models.AdvResponse, error) {
	if err := s.repo.UpdateAdv(ctx, id, adv); err != nil {
		return nil, err
	}

	return s.repo.GetAdvByID(ctx, id)
}

func (s *Service) GetAdvsWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*models.AdvResponse, error) {
	key := s.generateCacheKey(limit, offset, posts, priority)
	
	advs, err := s.cache.GetAdv(ctx, key)
	if err != nil {
		if !errors.Is(err, repository.ErrKeyNotFound) {
			s.log.Error(err)
		}

		// идем в бд
		advs, err = s.repo.GetAdvsWithFilter(ctx, limit, offset, posts, priority)
		if err != nil {
			return nil, err
		}

		// сохраняем в кэш
		if err := s.cache.SetAdv(ctx, key, advs); err != nil {
			s.log.Error(err)
		}
	}

	return advs, nil
}

func(s *Service) DeleteAdvsCache(ctx context.Context) error {
	return s.cache.DeleteAdvsCache(ctx)
}

func (s *Service) generateCacheKey(limit, offset int, posts, priority []string) string {
	key := fmt.Sprintf("%d:%d:%s:%s", limit, offset, strings.Join(posts, ","), strings.Join(priority, ","))

	// Создаем хеш
	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return "advs:" + hash
}
