package router

import (
	td "snakealive/m/internal/gateway/trip/delivery"
	ud "snakealive/m/internal/gateway/user/delivery"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/middlewares/http"
	auth_service "snakealive/m/pkg/services/auth"
	trip_service "snakealive/m/pkg/services/trip"

	"github.com/fasthttp/router"
	"go.uber.org/zap"
)

type RouterConfig struct {
	AuthGRPC     auth_service.AuthServiceClient
	UserDelivery ud.UserDelivery

	TripGRPC            trip_service.TripServiceClient
	TripGatewayDelivery td.TripGatewayDelivery

	Logger *zap.Logger
}

func SetupRouter(cfg RouterConfig) (r *router.Router) {
	r = router.New()
	lgrMw := http.NewLoggingMiddleware(cfg.Logger)
	authMw := http.NewSessionValidatorMiddleware(
		cfg.AuthGRPC,
		error_adapter.NewGrpcToHttpAdapter(grpc_errors.PreparedAuthErrors, grpc_errors.CommonAuthError),
	)

	r.POST(cnst.LoginURL, lgrMw(cfg.UserDelivery.Login))
	r.DELETE(cnst.LogoutURL, lgrMw(authMw(cfg.UserDelivery.Logout)))
	r.GET(cnst.ProfileURL, lgrMw(authMw(cfg.UserDelivery.GetProfile)))
	r.PATCH(cnst.ProfileURL, lgrMw(authMw(cfg.UserDelivery.UpdateProfile)))
	r.POST(cnst.RegisterURL, lgrMw(cfg.UserDelivery.Register))

	r.GET(cnst.TripURL, lgrMw(authMw(cfg.TripGatewayDelivery.Trip)))
	r.POST(cnst.TripPostURL, lgrMw(authMw(cfg.TripGatewayDelivery.AddTrip)))
	r.PATCH(cnst.TripURL, lgrMw(authMw(cfg.TripGatewayDelivery.Update)))
	r.DELETE(cnst.TripURL, lgrMw(authMw(cfg.TripGatewayDelivery.Delete)))

	return
}
