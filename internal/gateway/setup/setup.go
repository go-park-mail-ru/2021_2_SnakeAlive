package setup

import (
	fsthp_router "github.com/fasthttp/router"
	"google.golang.org/grpc"
	"snakealive/m/internal/gateway/config"
	"snakealive/m/internal/gateway/router"
	"snakealive/m/internal/gateway/user/delivery"
	"snakealive/m/internal/gateway/user/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	auth_service "snakealive/m/pkg/services/auth"
)

func Setup(cfg config.Config) (r *fsthp_router.Router, stopFunc func(), err error) {
	conn, err := grpc.Dial(cfg.AuthServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return r, stopFunc, err
	}

	userGRPC := auth_service.NewAuthServiceClient(conn)
	userUsecase := usecase.NewUserUsecase(userGRPC)
	userDelivery := delivery.NewUserDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		userUsecase,
	)
	r = router.SetupRouter(router.RouterConfig{
		AuthGRPC:     userGRPC,
		UserDelivery: userDelivery,
		Logger:       cfg.Logger,
	})

	stopFunc = func() {
		conn.Close()
	}

	return
}
