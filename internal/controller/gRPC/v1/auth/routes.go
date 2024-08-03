package auth_router

import (
	auth_model "auth/internal/domain/auth/model"
	"context"
	authv1 "github.com/winterkl/auth_protobuf/gen/go/proto/auth"
	"google.golang.org/protobuf/types/known/emptypb"
)

type authRoute struct {
	authv1.UnimplementedAuthServer
	auth UseCase
}

func New(auth UseCase) *authRoute {
	return &authRoute{
		auth: auth,
	}
}

func (r *authRoute) Register(ctx context.Context, in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	//Конвертируем gRPC-модель в domain-модель
	model := NewMapper(in).ToInner().(auth_model.CreateUserRequest)

	//Бизнес логика
	response, err := r.auth.CreateUser(ctx, model)
	if err != nil {
		return nil, handleError(err)
	}

	//Конвертируем domain-модель в gRPC-модель и возвращаем её\
	return NewMapper(response).ToOuter().(*authv1.RegisterResponse), nil

}
func (r *authRoute) GetToken(ctx context.Context, in *authv1.GetTokenRequest) (*authv1.GetTokenResponse, error) {
	//Конвертируем gRPC-модель в domain-модель
	model := NewMapper(in).ToInner().(auth_model.GetTokenRequest)

	//Бизнес логика
	response, err := r.auth.GetToken(ctx, model)
	if err != nil {
		return nil, handleError(err)
	}

	//Конвертируем domain-модель в gRPC-модель и возвращаем её
	return NewMapper(response).ToOuter().(*authv1.GetTokenResponse), nil
}
func (r *authRoute) ValidateToken(ctx context.Context, in *authv1.ValidateTokenRequest) (*emptypb.Empty, error) {
	//Конвертируем gRPC-модель в domain-модель
	model := NewMapper(in).ToInner().(auth_model.ValidateTokenRequest)

	//Бизнес логика
	if err := r.auth.ValidateToken(ctx, model); err != nil {
		return nil, handleError(err)
	}

	return nil, nil
}
