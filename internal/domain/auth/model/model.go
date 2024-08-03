package auth_model

import (
	"auth/internal/app_errors"
	auth_entity "auth/internal/domain/auth/entity"
	"auth/pkg/hash"
)

type CreateUserRequest struct {
	Login    string
	Password string
}

func (m *CreateUserRequest) GetEntity() auth_entity.User {
	return auth_entity.User{
		Login:    m.Login,
		Password: hash.GetMD5(m.Password),
	}
}

type CreateUserResponse struct {
	ID int
}

func NewGetUserResponse(id int) CreateUserResponse {
	return CreateUserResponse{
		ID: id,
	}
}

type GetTokenRequest struct {
	Login    string
	Password string
}

func (m *GetTokenRequest) Validate() error {
	if m.Login == "" {
		return &app_errors.IsRequired{Field: "login"}
	}
	if m.Password == "" {
		return &app_errors.IsRequired{Field: "password"}
	}
	return nil
}

func (m *GetTokenRequest) GetEntity() auth_entity.User {
	return auth_entity.User{
		Login:    m.Login,
		Password: hash.GetMD5(m.Password),
	}
}

type GetTokenResponse struct {
	Token string
}

func NewGetTokenResponse(token string) GetTokenResponse {
	return GetTokenResponse{
		Token: token,
	}
}

type ValidateTokenRequest struct {
	Token string
}

func (m *ValidateTokenRequest) Validate() error {
	if m.Token == "" {
		return &app_errors.IsRequired{Field: "token"}
	}
	return nil
}
