package countryDelivery

import (
	"snakealive/m/internal/domain"
	"strconv"

	cr "snakealive/m/internal/country/repository"
	cu "snakealive/m/internal/country/usecase"
	"snakealive/m/internal/entities"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type counryHandler struct {
	CountryUseCase domain.CountryUseCase
}

func NewCountryHandler(CountryUseCase domain.CountryUseCase) domain.CountryHandler {
	return &counryHandler{
		CountryUseCase: CountryUseCase,
	}
}

func CreateDelivery(db *pgxpool.Pool, l *logs.Logger) domain.CountryHandler {
	countryLayer := NewCountryHandler(cu.NewCountryUseCase(cr.NewCountryStorage(db), l))
	return countryLayer
}

func SetUpCountryRouter(db *pgxpool.Pool, r *router.Router, l *logs.Logger) *router.Router {
	countryHandler := CreateDelivery(db, l)
	r.GET(cnst.CountryListURL, logs.AccessLogMiddleware(l, countryHandler.GetCountriesList))
	r.GET(cnst.CountryIdURL, logs.AccessLogMiddleware(l, countryHandler.GetById))
	r.GET(cnst.CountryNameURL, logs.AccessLogMiddleware(l, countryHandler.GetByName))
	return r
}

func (s *counryHandler) GetCountriesList(ctx *fasthttp.RequestCtx) {
	code, bytes := s.CountryUseCase.GetCountriesList()

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}

func (s *counryHandler) GetByName(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("name").(string)
	trans := entities.CountryTrans[param]
	code, bytes := s.CountryUseCase.GetByName(trans)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}

func (s *counryHandler) GetById(ctx *fasthttp.RequestCtx) {
	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	code, bytes := s.CountryUseCase.GetById(param)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}
