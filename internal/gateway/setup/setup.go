package setup

import (
	"snakealive/m/internal/gateway/config"
	country_delivery "snakealive/m/internal/gateway/country/delivery"
	"snakealive/m/internal/gateway/country/repository"
	"snakealive/m/internal/gateway/country/usecase"
	media_delivery "snakealive/m/internal/gateway/media/delivery"
	media_usecase "snakealive/m/internal/gateway/media/usecase"
	review_delivery "snakealive/m/internal/gateway/review/delivery"
	review_usecase "snakealive/m/internal/gateway/review/usecase"
	"snakealive/m/internal/gateway/router"
	"snakealive/m/internal/gateway/sight/delivery"
	sight_usecase "snakealive/m/internal/gateway/sight/usecase"
	td "snakealive/m/internal/gateway/trip/delivery"
	tu "snakealive/m/internal/gateway/trip/usecase"
	ud "snakealive/m/internal/gateway/user/delivery"
	uu "snakealive/m/internal/gateway/user/usecase"
	"snakealive/m/pkg/error_adapter"
	fasthttpprom "snakealive/m/pkg/fasthttp_prom"
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/hasher"
	"snakealive/m/pkg/helpers"
	auth_service "snakealive/m/pkg/services/auth"
	review_service "snakealive/m/pkg/services/review"
	sight_service "snakealive/m/pkg/services/sight"
	trip_service "snakealive/m/pkg/services/trip"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"google.golang.org/grpc"
)

func Setup(cfg config.Config) (p fasthttpprom.Router, stopFunc func(), err error) {
	pgxConn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return  p, stopFunc, err
	}

	countryRepo := repository.NewLoggingMiddleware(
		cfg.Logger.Sugar(),
		repository.NewCountryStorage(repository.NewQueryFactory(), pgxConn),
	)
	countryUsecase := usecase.NewCountryUsecase(countryRepo)
	countryDelivery := country_delivery.NewCountryDelivery(countryUsecase, error_adapter.NewErrorToHttpAdapter(
		grpc_errors.PreparedCountryErrors, grpc_errors.CommonError,
	))

	conn, err := grpc.Dial(cfg.AuthServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return  p, stopFunc, err
	}
	userGRPC := auth_service.NewAuthServiceClient(conn)
	userUsecase := uu.NewUserUsecase(userGRPC)
	userDelivery := ud.NewUserDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		userUsecase,
	)

	sightConn, err := grpc.Dial(cfg.SightServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return p, stopFunc, err
	}
	sightGRPC := sight_service.NewSightServiceClient(sightConn)
	sightUsecase := sight_usecase.NewSightGatewayUseCase(sightGRPC)
	sightDelivery := delivery.NewSightDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		sightUsecase,
	)

	tripConn, err := grpc.Dial(cfg.TripServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return  p, stopFunc, err
	}
	tripGRPC := trip_service.NewTripServiceClient(tripConn)
	tripGatewayUseCase := tu.NewTripGatewayUseCase(tripGRPC, sightGRPC, userGRPC)
	tripDelivery := td.NewTripGetewayDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		tripGatewayUseCase,
	)

	reviewConn, err := grpc.Dial(cfg.ReviewServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return  p, stopFunc, err
	}
	reviewGRPC := review_service.NewReviewServiceClient(reviewConn)
	reviewUsecase := review_usecase.NewReviewGatewayUseCase(reviewGRPC)
	reviewDelivery := review_delivery.NewReviewGatewayDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		reviewUsecase,
	)

	s3Client := s3.New(
		session.Must(session.NewSession()),
		aws.NewConfig().WithEndpoint(cfg.S3Endpoint),
		aws.NewConfig().WithRegion("ru-msk"),
		aws.NewConfig().WithCredentials(
			credentials.NewStaticCredentials(
				cfg.ID,
				cfg.SecretKey,
				"",
			),
		),
	)
	mediaUsecase := media_usecase.NewMediaUsecase(
		s3Client, hasher.NewHasher(5),
		cfg.DefaultBucket, cfg.S3PublicEndpoint,
	)
	mediaDelivery := media_delivery.NewMediaDelivery(mediaUsecase,
		error_adapter.NewErrorToHttpAdapter(
			grpc_errors.PreparedCountryErrors, grpc_errors.CommonError,
		))

	p = router.SetupRouter(router.RouterConfig{
		AuthGRPC:            userGRPC,
		UserDelivery:        userDelivery,
		TripGatewayDelivery: tripDelivery,
		SightDelivery:       sightDelivery,
		ReviewDelivery:      reviewDelivery,
		CountryDelivery:     countryDelivery,
		MediaDelivery:       mediaDelivery,
		Logger:              cfg.Logger,
	})

	stopFunc = func() {
		conn.Close()
		tripConn.Close()
		sightConn.Close()
		pgxConn.Close()
	}

	return
}
