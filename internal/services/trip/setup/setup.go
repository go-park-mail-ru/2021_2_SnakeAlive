package setup

import (
	"snakealive/m/internal/services/trip/config"
	"snakealive/m/internal/services/trip/delivery"
	"snakealive/m/internal/services/trip/repository"
	trip_usecase "snakealive/m/internal/services/trip/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/helpers"
	trip_service "snakealive/m/pkg/services/trip"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	zap_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func SetupServer(cfg config.Config) (server *grpc.Server, cancel func(), err error) {
	conn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return server, cancel, err
	}

	tripRepo := repository.NewTripRepository(conn)
	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := delivery.NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
		),
	)
	trip_service.RegisterTripServiceServer(server, tripDelivery)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}
