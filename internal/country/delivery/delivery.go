package countryDelivery

import (
	"snakealive/m/pkg/domain"
	"strconv"

	cr "snakealive/m/internal/country/repository"
	cu "snakealive/m/internal/country/usecase"
	"snakealive/m/internal/entities"
	logs "snakealive/m/internal/logger"
	cnst "snakealive/m/pkg/constants"

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
	r.GET(cnst.CountryListURL, countryHandler.GetCountriesList)
	r.GET(cnst.CountryIdURL, countryHandler.GetById)
	r.GET(cnst.CountryNameURL, countryHandler.GetByName)
	return r
}

func (s *counryHandler) GetCountriesList(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	code, bytes := s.CountryUseCase.GetCountriesList()

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *counryHandler) GetByName(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	param, _ := ctx.UserValue("name").(string)
	trans := entities.CountryTrans[param]
	code, bytes := s.CountryUseCase.GetByName(trans)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *counryHandler) GetById(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		logger.Error("error while getting sid param: ", zap.Error(err))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	code, bytes := s.CountryUseCase.GetById(param)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}
