package placeDelivery

import (
	"snakealive/m/pkg/domain"

	pr "snakealive/m/internal/place/repository"
	pu "snakealive/m/internal/place/usecase"
	cnst "snakealive/m/pkg/constants"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

const CookieName = "SnakeAlive"

type PlaceHandler interface {
	PlacesByCountry(ctx *fasthttp.RequestCtx)
	Place(ctx *fasthttp.RequestCtx)
}

type placeHandler struct {
	PlaceUseCase domain.PlaceUseCase
}

func NewPlaceHandler(PlaceUseCase domain.PlaceUseCase) PlaceHandler {
	return &placeHandler{
		PlaceUseCase: PlaceUseCase,
	}
}

func CreateDelivery(db *pgxpool.Pool) PlaceHandler {
	placeLayer := NewPlaceHandler(pu.NewPlaceUseCase(pr.NewPlaceStorage(db)))
	return placeLayer
}

func SetUpPlaceRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	placeHandler := CreateDelivery(db)
	r.GET(cnst.CountryURL, placeHandler.PlacesByCountry)
	r.GET(cnst.SightURL, placeHandler.Place)
	return r
}

func (s *placeHandler) PlacesByCountry(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("name").(string)
	code, bytes := s.PlaceUseCase.GetPlaceListByName(param)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}
