package grpc_errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snakealive/m/pkg/errors"
)

var (
	PreparedAuthServiceErrorMap = map[error]error{
		errors.UserDoesNotExist:    status.Error(codes.NotFound, "user not found"),
		errors.SessionDoesNotExist: status.Error(codes.NotFound, "session not found"),
		errors.WrongUserPassword:   status.Error(codes.PermissionDenied, "wrong password"),
	}
)
