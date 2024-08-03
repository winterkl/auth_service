package v1

import (
	auth_router "auth/internal/controller/gRPC/v1/auth"
	"auth/pkg/grpc_server"
	authv1 "github.com/winterkl/auth_protobuf/gen/go/proto/auth"
)

type UseCase struct {
	Auth auth_router.UseCase
}

func Register(gRPCServer *grpc_server.GRPCServer, useCase UseCase) {
	authv1.RegisterAuthServer(gRPCServer, auth_router.New(useCase.Auth))
}
