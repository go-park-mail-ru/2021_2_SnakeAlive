package placeDelivery

import (
	"snakealive/m/pkg/domain"
	"strconv"

	"snakealive/m/internal/entities"
	logs "snakealive/m/internal/logger"
	pr "snakealive/m/internal/place/repository"
	pu "snakealive/m/internal/place/usecase"
	cnst "snakealive/m/pkg/constants"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

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
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

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
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *placeHandler) Place(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))

	sight, err := s.PlaceUseCase.GetById(param)
	if err != nil {
		logger.Error("error while getting sight: ", zap.Error(err))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	code, bytes := s.PlaceUseCase.GetSight(sight)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}
