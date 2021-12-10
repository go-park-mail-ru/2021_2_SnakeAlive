package setup

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	zap_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"snakealive/m/internal/services/auth/config"
	"snakealive/m/internal/services/auth/delivery"
	"snakealive/m/internal/services/auth/repository"
	auth_usecase "snakealive/m/internal/services/auth/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/hasher"
	"snakealive/m/pkg/helpers"
	auth_service "snakealive/m/pkg/services/auth"
)

func SetupServer(cfg config.Config) (server *grpc.Server, cancel func(), err error) {
	conn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return server, cancel, err
	}

	authRepo := repository.NewLoggingMiddleware(
		cfg.Logger.Sugar(),
		repository.NewAuthRepository(
			repository.NewQueryFactory(), conn,
		),
	)
	authUsecase := auth_usecase.NewAuthUseCase(hasher.NewHasher(cfg.PrefixLen), authRepo)
	authDelivery := delivery.NewAuthDelivery(authUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	auth_service.RegisterAuthServiceServer(server, authDelivery)
	grpc_prometheus.Register(server)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}
