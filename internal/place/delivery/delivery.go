package placeDelivery

import (
	"snakealive/m/pkg/domain"
	"strconv"

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
	/*param, _ := ctx.UserValue("name").(string)
	if _, found := s.PlaceUseCase.Get(param); !found {
		log.Printf("country not found")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	response, _ := s.PlaceUseCase.Get(param)
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)*/
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
