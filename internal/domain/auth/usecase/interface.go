package auth_usecase

import (
	auth_entity "auth/internal/domain/auth/entity"
	"context"
	"time"
)

type Repository interface {
	CreateUser(ctx context.Context, user auth_entity.User) (userID int, err error)
	GetUserByAuthData(ctx context.Context, user auth_entity.User) error
}

type CacheRepository interface {
	SetString(ctx context.Context, key, value string, exp time.Duration) error
	GetString(ctx context.Context, key string) (value string, err error)
}

type JwtAuth interface {
	GenerateToken(login string) (string, error)
	ParseToken(tokenStr string) (string, error)
}
