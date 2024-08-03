package cache_repo

import (
	"auth/internal/app_errors"
	"auth/pkg/redis_db"
	"context"
	"fmt"
	"time"
)

type CacheRepository struct {
	db *redis_db.Client
}

func New(db *redis_db.Client) *CacheRepository {
	return &CacheRepository{
		db: db,
	}
}

func (r *CacheRepository) SetString(ctx context.Context, key, token string, ttl time.Duration) error {
	if err := r.db.Set(ctx, key, token, ttl).Err(); err != nil {
		return fmt.Errorf("Set: %w", err)
	}
	return nil

	//TODO: в key можно зашивать appID
}

func (r *CacheRepository) GetString(ctx context.Context, key string) (string, error) {
	token, err := r.db.Get(ctx, key).Result()
	if err != nil {
		if token == "" {
			return "", &app_errors.TokenNotFound{}
		}
		return "", fmt.Errorf("Get: %w", err)
	}

	return token, nil
}
