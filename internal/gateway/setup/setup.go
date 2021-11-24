package setup

import (
	"snakealive/m/internal/gateway/config"
	country_delivery "snakealive/m/internal/gateway/country/delivery"
	"snakealive/m/internal/gateway/country/repository"
	"snakealive/m/internal/gateway/country/usecase"
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
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/helpers"
	auth_service "snakealive/m/pkg/services/auth"
	review_service "snakealive/m/pkg/services/review"
	sight_service "snakealive/m/pkg/services/sight"
	trip_service "snakealive/m/pkg/services/trip"

	fsthp_router "github.com/fasthttp/router"
	"google.golang.org/grpc"
)

func Setup(cfg config.Config) (r *fsthp_router.Router, stopFunc func(), err error) {
	pgxConn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return r, stopFunc, err
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
		return r, stopFunc, err
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
		return r, stopFunc, err
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
		return r, stopFunc, err
	}
	tripGRPC := trip_service.NewTripServiceClient(tripConn)
	tripGatewayUseCase := tu.NewTripGatewayUseCase(tripGRPC, sightGRPC)
	tripDelivery := td.NewTripGetewayDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		tripGatewayUseCase,
	)

	reviewConn, err := grpc.Dial(cfg.ReviewServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return r, stopFunc, err
	}
	reviewGRPC := review_service.NewReviewServiceClient(reviewConn)
	reviewUsecase := review_usecase.NewReviewGatewayUseCase(reviewGRPC)
	reviewDelivery := review_delivery.NewReviewGatewayDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		reviewUsecase,
	)

	r = router.SetupRouter(router.RouterConfig{
		AuthGRPC:            userGRPC,
		UserDelivery:        userDelivery,
		TripGatewayDelivery: tripDelivery,
		SightDelivery:       sightDelivery,
		ReviewDelivery:      reviewDelivery,
		CountryDelivery:     countryDelivery,
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
