package auth_router

import (
	auth_model "auth/internal/domain/auth/model"
	"context"
)

type UseCase interface {
	CreateUser(ctx context.Context, model auth_model.CreateUserRequest) (auth_model.CreateUserResponse, error)
	GetToken(ctx context.Context, model auth_model.GetTokenRequest) (auth_model.GetTokenResponse, error)
	ValidateToken(ctx context.Context, model auth_model.ValidateTokenRequest) error
}
