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
	"go.uber.org/zap"
)

type counryHandler struct {
	CountryUseCase domain.CountryUseCase
}

func NewCountryHandler(CountryUseCase domain.CountryUseCase) domain.CountryHandler {
	return &counryHandler{
		CountryUseCase: CountryUseCase,
	}
}

func CreateDelivery(db *pgxpool.Pool) domain.CountryHandler {
	countryLayer := NewCountryHandler(cu.NewCountryUseCase(cr.NewCountryStorage(db)))
	return countryLayer
}

func SetUpCountryRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	countryHandler := CreateDelivery(db)
	r.GET(cnst.CountryListURL, logs.AccessLogMiddleware(countryHandler.GetCountriesList))
	r.GET(cnst.CountryIdURL, logs.AccessLogMiddleware(countryHandler.GetById))
	r.GET(cnst.CountryNameURL, logs.AccessLogMiddleware(countryHandler.GetByName))
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
	logger := logs.GetLogger()

	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		logger.Error("error while getting sid param: ", zap.Error(err))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	code, bytes := s.CountryUseCase.GetById(param)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}
