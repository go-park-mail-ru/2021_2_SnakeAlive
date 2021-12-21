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
	fasthttpprom "snakealive/m/pkg/fasthttp_prom"
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

func SetupRouter(cfg RouterConfig) (p fasthttpprom.Router) {
	p = fasthttpprom.NewRouter("gateway")
	p.Use(router.New())

	lgrMw := http.NewLoggingMiddleware(cfg.Logger)
	authMw := http.NewSessionValidatorMiddleware(
		cfg.AuthGRPC,
		error_adapter.NewGrpcToHttpAdapter(grpc_errors.PreparedAuthErrors, grpc_errors.CommonAuthError),
	)

	p.POST(cnst.LoginURL, lgrMw(cfg.UserDelivery.Login))
	p.DELETE(cnst.LogoutURL, lgrMw(authMw(cfg.UserDelivery.Logout)))
	p.GET(cnst.ProfileURL, lgrMw(authMw(cfg.UserDelivery.GetProfile)))
	p.PATCH(cnst.ProfileURL, lgrMw(authMw(cfg.UserDelivery.UpdateProfile)))
	p.POST(cnst.RegisterURL, lgrMw(cfg.UserDelivery.Register))
	p.GET(cnst.UserInfoURL, lgrMw(cfg.UserDelivery.GetUserInfo))

	p.GET(cnst.TripURL, lgrMw(authMw(cfg.TripGatewayDelivery.Trip)))
	p.POST(cnst.TripPostURL, lgrMw(authMw(cfg.TripGatewayDelivery.AddTrip)))
	p.PATCH(cnst.TripURL, lgrMw(authMw(cfg.TripGatewayDelivery.UpdateTrip)))

	p.DELETE(cnst.TripURL, lgrMw(authMw(cfg.TripGatewayDelivery.DeleteTrip)))
	p.GET(cnst.AlbumURL, lgrMw(authMw(cfg.TripGatewayDelivery.Album)))
	p.POST(cnst.AlbumAddURL, lgrMw(authMw(cfg.TripGatewayDelivery.AddAlbum)))
	p.PATCH(cnst.AlbumURL, lgrMw(authMw(cfg.TripGatewayDelivery.UpdateAlbum)))
	p.DELETE(cnst.AlbumURL, lgrMw(authMw(cfg.TripGatewayDelivery.DeleteAlbum)))
	p.GET(cnst.SightsByTripURL, lgrMw(cfg.TripGatewayDelivery.SightsByTrip))
	p.GET(cnst.TripsByUserURL, lgrMw(authMw(cfg.TripGatewayDelivery.TripsByUser)))
	p.GET(cnst.AlbumsByUserURL, lgrMw(authMw(cfg.TripGatewayDelivery.AlbumsByUser)))
	p.POST(cnst.AddTripUserURL, lgrMw(authMw(cfg.TripGatewayDelivery.AddTripUser)))
	p.POST(cnst.ShareTripURL, lgrMw(authMw(cfg.TripGatewayDelivery.ShareLink)))
	p.GET(cnst.SharedTripURL, lgrMw(authMw(cfg.TripGatewayDelivery.AddUserByLink)))

	p.GET(cnst.SightsByCountryURL, lgrMw(cfg.SightDelivery.GetSightByCountry))
	p.GET(cnst.SightURL, lgrMw(cfg.SightDelivery.GetSightByID))
	p.PATCH(cnst.SightSearch, lgrMw(cfg.SightDelivery.SearchSights))
	p.GET(cnst.SightTag, lgrMw(cfg.SightDelivery.GetSightByTag))
	p.GET(cnst.Tags, lgrMw(cfg.SightDelivery.GetTags))
	p.GET(cnst.RandomTags, lgrMw(cfg.SightDelivery.GetRandomTags))

	p.POST(cnst.ReviewAddURL, lgrMw(authMw(cfg.ReviewDelivery.AddReviewToPlace)))
	p.GET(cnst.ReviewURL, lgrMw(cfg.ReviewDelivery.ReviewsByPlace))
	p.DELETE(cnst.ReviewURL, lgrMw(authMw(cfg.ReviewDelivery.DelReview)))

	p.GET(cnst.CountryNameURL, lgrMw(cfg.CountryDelivery.GetCountryByName))
	p.GET(cnst.CountryIdURL, lgrMw(cfg.CountryDelivery.GetCountryByID))
	p.GET(cnst.CountryListURL, lgrMw(cfg.CountryDelivery.ListCountries))
	p.POST(cnst.UploadURL, lgrMw(cfg.MediaDelivery.UploadFile))

	return
}
