package placeDelivery

import (
	"snakealive/m/internal/domain"
	"strconv"

	"snakealive/m/internal/entities"
	pr "snakealive/m/internal/place/repository"
	pu "snakealive/m/internal/place/usecase"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type placeHandler struct {
	PlaceUseCase domain.PlaceUseCase
}

func NewPlaceHandler(PlaceUseCase domain.PlaceUseCase) domain.PlaceHandler {
	return &placeHandler{
		PlaceUseCase: PlaceUseCase,
	}
}

func CreateDelivery(db *pgxpool.Pool, l *logs.Logger) domain.PlaceHandler {
	placeLayer := NewPlaceHandler(pu.NewPlaceUseCase(pr.NewPlaceStorage(db), l))
	return placeLayer
}

func SetUpPlaceRouter(db *pgxpool.Pool, r *router.Router, l *logs.Logger) *router.Router {
	placeHandler := CreateDelivery(db, l)
	r.GET(cnst.SightsByCountruURL, logs.AccessLogMiddleware(l, placeHandler.PlacesByCountry))
	r.GET(cnst.SightURL, logs.AccessLogMiddleware(l, placeHandler.Place))
	return r
}

func (s *placeHandler) PlacesByCountry(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("name").(string)
	trans := entities.CountryTrans[param]
	bytes, err := s.PlaceUseCase.GetPlacesByCountry(trans)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *placeHandler) Place(ctx *fasthttp.RequestCtx) {
	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	sight, err := s.PlaceUseCase.GetById(param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	code, bytes := s.PlaceUseCase.GetSight(sight)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}
