package auth_usecase

import (
	"auth/internal/app_errors"
	auth_model "auth/internal/domain/auth/model"
	"context"
	"errors"
	"fmt"
	"time"
)

const tokenTTL = time.Hour * 12

type UseCase struct {
	repo  Repository
	cache CacheRepository
	jwt   JwtAuth
}

func New(repo Repository, cache CacheRepository, jwt JwtAuth) *UseCase {
	return &UseCase{
		repo:  repo,
		cache: cache,
		jwt:   jwt,
	}
}

func (uc *UseCase) CreateUser(ctx context.Context, model auth_model.CreateUserRequest) (auth_model.CreateUserResponse, error) {
	user := model.GetEntity()
	userID, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return auth_model.CreateUserResponse{}, fmt.Errorf("auth_usecase --> CreateUser --> repo.CreateUser: %w", err)
	}
	return auth_model.NewGetUserResponse(userID), err
}

func (uc *UseCase) GetToken(ctx context.Context, model auth_model.GetTokenRequest) (auth_model.GetTokenResponse, error) {
	if err := model.Validate(); err != nil {
		return auth_model.GetTokenResponse{}, fmt.Errorf("auth_usecase --> GetToken --> model.Validate: %w", err)
	}

	user := model.GetEntity()
	if err := uc.repo.GetUserByAuthData(ctx, user); err != nil {
		var userNotFound *app_errors.UserNotFound
		if errors.As(err, &userNotFound) {
			return auth_model.GetTokenResponse{}, &app_errors.IncorrectAuthData{}
		}
		return auth_model.GetTokenResponse{}, fmt.Errorf("auth_usecase --> GetString --> repo.GetString: %w", err)
	}

	token, err := uc.jwt.GenerateToken(model.Login)
	if err != nil {
		return auth_model.GetTokenResponse{}, fmt.Errorf("auth_usecase --> GetUserByAuthData --> jwt.GenerateToken: %w", err)
	}

	if err = uc.cache.SetString(ctx, user.Login, token, tokenTTL); err != nil {
		return auth_model.GetTokenResponse{}, fmt.Errorf("auth_usecase --> SetString --> cache.SetString --> %w", err)
	}

	return auth_model.NewGetTokenResponse(token), nil
}

func (uc *UseCase) ValidateToken(ctx context.Context, model auth_model.ValidateTokenRequest) error {
	if err := model.Validate(); err != nil {
		return fmt.Errorf("auth_usecase --> ValidateToken --> model.Validate: %w", err)
	}

	login, err := uc.jwt.ParseToken(model.Token)
	if err != nil {
		return fmt.Errorf("auth_usecase --> ValidateToken --> jwt.ParseToken --> %w", err)
	}

	cacheToken, err := uc.cache.GetString(ctx, login)
	if err != nil {
		return fmt.Errorf("auth_usecase --> GetString --> cache.GetString --> %w", err)
	}

	if cacheToken != model.Token {
		return &app_errors.InvalidValidateToken{}
	}

	return nil
}
