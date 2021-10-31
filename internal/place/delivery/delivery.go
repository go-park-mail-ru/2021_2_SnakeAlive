package placeDelivery

import (
	"log"
	"snakealive/m/pkg/domain"
	"strconv"

	"snakealive/m/internal/entities"
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

	trans := entities.CountryTrans[param]
	bytes, err := s.PlaceUseCase.GetPlacesByCountry(trans)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf("error while getting list: %s", err)
		ctx.Write([]byte("{}"))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *placeHandler) Place(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))

	sight, err := s.PlaceUseCase.GetById(param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	code, bytes := s.PlaceUseCase.GetSight(sight)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}
