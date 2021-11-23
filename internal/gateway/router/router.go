package router

import (
	"github.com/fasthttp/router"
	"go.uber.org/zap"
	sight_delivery "snakealive/m/internal/gateway/sight/delivery"
	td "snakealive/m/internal/gateway/trip/delivery"
	ud "snakealive/m/internal/gateway/user/delivery"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/middlewares/http"
	auth_service "snakealive/m/pkg/services/auth"
)

type RouterConfig struct {
	AuthGRPC auth_service.AuthServiceClient

	UserDelivery        ud.UserDelivery
	TripGatewayDelivery td.TripGatewayDelivery
	SightDelivery       sight_delivery.SightDelivery

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

	r.GET(cnst.SightsByCountryURL, lgrMw(cfg.SightDelivery.GetSightByCountry))
	r.GET(cnst.SightURL, lgrMw(cfg.SightDelivery.GetSightByID))

	return
}
