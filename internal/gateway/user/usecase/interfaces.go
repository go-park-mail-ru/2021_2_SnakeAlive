package usecase

import (
	"context"

	auth_service "snakealive/m/pkg/services/auth"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type authGRPC interface {
	GetUserInfo(ctx context.Context, in *auth_service.GetUserRequest, opts ...grpc.CallOption) (*auth_service.UserInfo, error)
	LogoutUser(ctx context.Context, in *auth_service.Session, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LoginUser(ctx context.Context, in *auth_service.LoginRequest, opts ...grpc.CallOption) (*auth_service.LoginResponse, error)
	RegisterUser(ctx context.Context, in *auth_service.RegisterRequest, opts ...grpc.CallOption) (*auth_service.LoginResponse, error)
	GetUser(ctx context.Context, in *auth_service.GetUserRequest, opts ...grpc.CallOption) (*auth_service.GetUserResponse, error)
	UpdateUser(ctx context.Context, in *auth_service.UpdateUserRequest, opts ...grpc.CallOption) (*auth_service.GetUserResponse, error)
}
