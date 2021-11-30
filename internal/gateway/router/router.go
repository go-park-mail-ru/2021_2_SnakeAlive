package router

import (
	country_delivery "snakealive/m/internal/gateway/country/delivery"
	media_delivery "snakealive/m/internal/gateway/media/delivery"
	review_delivery "snakealive/m/internal/gateway/review/delivery"
	sight_delivery "snakealive/m/internal/gateway/sight/delivery"
	td "snakealive/m/internal/gateway/trip/delivery"
	ud "snakealive/m/internal/gateway/user/delivery"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/middlewares/http"
	auth_service "snakealive/m/pkg/services/auth"

	"github.com/fasthttp/router"
	"go.uber.org/zap"
)

type RouterConfig struct {
	AuthGRPC auth_service.AuthServiceClient

	UserDelivery        ud.UserDelivery
	TripGatewayDelivery td.TripGatewayDelivery
	SightDelivery       sight_delivery.SightDelivery
	ReviewDelivery      review_delivery.ReviewGatewayDelivery
	CountryDelivery     country_delivery.CountryDelivery
	MediaDelivery       media_delivery.MediaDelivery

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
	r.GET(cnst.UserInfoURL, lgrMw(cfg.UserDelivery.GetUserInfo))

	r.GET(cnst.TripURL, lgrMw(authMw(cfg.TripGatewayDelivery.Trip)))
	r.POST(cnst.TripPostURL, lgrMw(authMw(cfg.TripGatewayDelivery.AddTrip)))
	r.PATCH(cnst.TripURL, lgrMw(authMw(cfg.TripGatewayDelivery.UpdateTrip)))
	r.DELETE(cnst.TripURL, lgrMw(authMw(cfg.TripGatewayDelivery.DeleteTrip)))
	r.GET(cnst.AlbumURL, lgrMw(authMw(cfg.TripGatewayDelivery.Album)))
	r.POST(cnst.AlbumAddURL, lgrMw(authMw(cfg.TripGatewayDelivery.AddAlbum)))
	r.PATCH(cnst.AlbumURL, lgrMw(authMw(cfg.TripGatewayDelivery.UpdateAlbum)))
	r.DELETE(cnst.AlbumURL, lgrMw(authMw(cfg.TripGatewayDelivery.DeleteAlbum)))
	r.GET(cnst.SightsByTripURL, lgrMw(cfg.TripGatewayDelivery.SightsByTrip))
	r.GET(cnst.TripsByUserURL, lgrMw(authMw(cfg.TripGatewayDelivery.TripsByUser)))
	r.GET(cnst.AlbumsByUserURL, lgrMw(authMw(cfg.TripGatewayDelivery.AlbumsByUser)))

	r.GET(cnst.SightsByCountryURL, lgrMw(cfg.SightDelivery.GetSightByCountry))
	r.GET(cnst.SightURL, lgrMw(cfg.SightDelivery.GetSightByID))
	r.GET(cnst.SightSearch, lgrMw(cfg.SightDelivery.SearchSights))
	r.GET(cnst.SightTag, lgrMw(cfg.SightDelivery.GetSightByTag))

	r.POST(cnst.ReviewAddURL, lgrMw(authMw(cfg.ReviewDelivery.AddReviewToPlace)))
	r.GET(cnst.ReviewURL, lgrMw(cfg.ReviewDelivery.ReviewsByPlace))
	r.DELETE(cnst.ReviewURL, lgrMw(authMw(cfg.ReviewDelivery.DelReview)))

	r.GET(cnst.CountryNameURL, lgrMw(cfg.CountryDelivery.GetCountryByName))
	r.GET(cnst.CountryIdURL, lgrMw(cfg.CountryDelivery.GetCountryByID))
	r.GET(cnst.CountryListURL, lgrMw(cfg.CountryDelivery.ListCountries))

	r.POST(cnst.UploadURL, lgrMw(authMw(cfg.MediaDelivery.UploadFile)))

	return
}
