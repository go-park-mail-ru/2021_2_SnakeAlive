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
	"go.uber.org/zap"
)

type placeHandler struct {
	PlaceUseCase domain.PlaceUseCase
}

func NewPlaceHandler(PlaceUseCase domain.PlaceUseCase) domain.PlaceHandler {
	return &placeHandler{
		PlaceUseCase: PlaceUseCase,
	}
}

func CreateDelivery(db *pgxpool.Pool) domain.PlaceHandler {
	placeLayer := NewPlaceHandler(pu.NewPlaceUseCase(pr.NewPlaceStorage(db)))
	return placeLayer
}

func SetUpPlaceRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	placeHandler := CreateDelivery(db)
	r.GET(cnst.SightsByCountruURL, logs.AccessLogMiddleware(placeHandler.PlacesByCountry))
	r.GET(cnst.SightURL, logs.AccessLogMiddleware(placeHandler.Place))
	return r
}

func (s *placeHandler) PlacesByCountry(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()

	param, _ := ctx.UserValue("name").(string)
	trans := entities.CountryTrans[param]
	bytes, err := s.PlaceUseCase.GetPlacesByCountry(trans)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		logger.Error("error while getting list of places: ", zap.Error(err))
		ctx.Write([]byte("{}"))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *placeHandler) Place(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()

	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		logger.Error("error while getting sid param: ", zap.Error(err))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	sight, err := s.PlaceUseCase.GetById(param)
	if err != nil {
		logger.Error("error while getting sight: ", zap.Error(err))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	code, bytes := s.PlaceUseCase.GetSight(sight)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}
