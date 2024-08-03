package auth_router

import (
	"auth/internal/app_errors"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(err error) error {
	var errTokenNotFound *app_errors.TokenNotFound
	if errors.As(err, &errTokenNotFound) {
		return status.Error(codes.NotFound, errTokenNotFound.Error())
	}
	var errIsRequired *app_errors.IsRequired
	if errors.As(err, &errIsRequired) {
		return status.Error(codes.InvalidArgument, errIsRequired.Error())
	}
	var errInvalidValidateToken *app_errors.InvalidValidateToken
	if errors.As(err, &errInvalidValidateToken) {
		return status.Error(codes.InvalidArgument, errInvalidValidateToken.Error())
	}
	var errUserAlreadyExists *app_errors.UserAlreadyExists
	if errors.As(err, &errUserAlreadyExists) {
		return status.Error(codes.AlreadyExists, errUserAlreadyExists.Error())
	}
	var errIncorrectAuthData *app_errors.IncorrectAuthData
	if errors.As(err, &errIncorrectAuthData) {
		return status.Error(codes.InvalidArgument, errIncorrectAuthData.Error())
	}
	return status.Error(codes.Internal, err.Error())
}
