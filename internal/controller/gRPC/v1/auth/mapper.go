package auth_router

import (
	auth_model "auth/internal/domain/auth/model"
	authv1 "github.com/winterkl/auth_protobuf/gen/go/proto/auth"
)

type Mapper struct {
	object interface{}
}

func NewMapper(object interface{}) *Mapper {
	return &Mapper{object: object}
}

func (m *Mapper) ToInner() interface{} {
	switch object := m.object.(type) {
	case *authv1.RegisterRequest:
		return m.createRegisterRequestToInner(object)
	case *authv1.GetTokenRequest:
		return m.createGetTokenRequestToInner(object)
	case *authv1.ValidateTokenRequest:
		return m.createValidateTokenRequestToInner(object)
	}
	return nil
}

func (m *Mapper) createRegisterRequestToInner(object *authv1.RegisterRequest) auth_model.CreateUserRequest {
	return auth_model.CreateUserRequest{
		Login:    object.Login,
		Password: object.Password,
	}
}

func (m *Mapper) createGetTokenRequestToInner(object *authv1.GetTokenRequest) auth_model.GetTokenRequest {
	return auth_model.GetTokenRequest{
		Login:    object.Login,
		Password: object.Password,
	}
}

func (m *Mapper) createValidateTokenRequestToInner(object *authv1.ValidateTokenRequest) auth_model.ValidateTokenRequest {
	return auth_model.ValidateTokenRequest{
		Token: object.Token,
	}
}

func (m *Mapper) ToOuter() interface{} {
	switch object := m.object.(type) {
	case auth_model.CreateUserResponse:
		return m.createRegisterResponseToOuter(object)
	case auth_model.GetTokenResponse:
		return m.createGetTokenResponseToOuter(object)
	}
	return nil
}

func (m *Mapper) createRegisterResponseToOuter(object auth_model.CreateUserResponse) *authv1.RegisterResponse {
	return &authv1.RegisterResponse{
		Id: int64(object.ID),
	}
}

func (m *Mapper) createGetTokenResponseToOuter(object auth_model.GetTokenResponse) *authv1.GetTokenResponse {
	return &authv1.GetTokenResponse{
		Token: object.Token,
	}
}
