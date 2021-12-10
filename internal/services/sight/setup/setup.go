package setup

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	zap_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"snakealive/m/internal/services/sight/config"
	"snakealive/m/internal/services/sight/delivery"
	"snakealive/m/internal/services/sight/repository"
	sight_usecase "snakealive/m/internal/services/sight/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/helpers"
	sight_service "snakealive/m/pkg/services/sight"
)

func SetupServer(cfg config.Config) (server *grpc.Server, cancel func(), err error) {
	conn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return server, cancel, err
	}

	sightRepo := repository.NewLoggingMiddleware(
		cfg.Logger.Sugar(),
		repository.NewSightRepository(
			repository.NewQueryFactory(), conn,
		),
	)
	sightUsecase := sight_usecase.NewSightUsecase(sightRepo)
	sightDelivery := delivery.NewSightDelivery(sightUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedSightServiceErrorMap))

	server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	sight_service.RegisterSightServiceServer(server, sightDelivery)
	grpc_prometheus.Register(server)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}
