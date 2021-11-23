package setup

import (
	"snakealive/m/internal/gateway/config"
	"snakealive/m/internal/gateway/router"
	"snakealive/m/internal/gateway/sight/delivery"
	sight_usecase "snakealive/m/internal/gateway/sight/usecase"
	td "snakealive/m/internal/gateway/trip/delivery"
	tu "snakealive/m/internal/gateway/trip/usecase"
	ud "snakealive/m/internal/gateway/user/delivery"
	uu "snakealive/m/internal/gateway/user/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	auth_service "snakealive/m/pkg/services/auth"
	sight_service "snakealive/m/pkg/services/sight"
	trip_service "snakealive/m/pkg/services/trip"

	fsthp_router "github.com/fasthttp/router"
	"google.golang.org/grpc"
)

func Setup(cfg config.Config) (r *fsthp_router.Router, stopFunc func(), err error) {
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

	tripConn, err := grpc.Dial(cfg.TripServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return r, stopFunc, err
	}
	tripGRPC := trip_service.NewTripServiceClient(tripConn)
	tripGatewayUseCase := tu.NewTripGatewayUseCase(tripGRPC)
	tripDelivery := td.NewTripGetewayDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		tripGatewayUseCase,
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

	r = router.SetupRouter(router.RouterConfig{
		AuthGRPC:            userGRPC,
		UserDelivery:        userDelivery,
		TripGatewayDelivery: tripDelivery,
		SightDelivery:       sightDelivery,
		Logger:              cfg.Logger,
	})

	stopFunc = func() {
		conn.Close()
		tripConn.Close()
		sightConn.Close()
	}

	return
}
