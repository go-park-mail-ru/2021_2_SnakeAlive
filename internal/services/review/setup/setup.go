package setup

import (
	"snakealive/m/internal/services/review/config"
	"snakealive/m/internal/services/review/delivery"
	"snakealive/m/internal/services/review/repository"
	review_usecase "snakealive/m/internal/services/review/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/helpers"
	review_service "snakealive/m/pkg/services/review"

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

	reviewRepo := repository.NewReviewRepository(conn)
	reviewUsecase := review_usecase.NewReviewUseCase(reviewRepo)
	reviewDelivery := delivery.NewReviewDelivery(reviewUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
		),
	)
	review_service.RegisterReviewServiceServer(server, reviewDelivery)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}
