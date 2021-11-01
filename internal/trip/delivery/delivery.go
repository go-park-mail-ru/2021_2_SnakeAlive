package tripDelivery

import (
	"encoding/json"
	"snakealive/m/pkg/domain"
	"strconv"

	cd "snakealive/m/internal/cookie/delivery"
	logs "snakealive/m/internal/logger"
	tr "snakealive/m/internal/trip/repository"
	tu "snakealive/m/internal/trip/usecase"
	cnst "snakealive/m/pkg/constants"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type TripHandler interface {
	Trip(ctx *fasthttp.RequestCtx)
	AddTrip(ctx *fasthttp.RequestCtx)
	Update(ctx *fasthttp.RequestCtx)
	Delete(ctx *fasthttp.RequestCtx)
}

type tripHandler struct {
	TripUseCase   domain.TripUseCase
	CookieHandler cd.CookieHandler
}

func NewTripHandler(TripUseCase domain.TripUseCase, CookieHandler cd.CookieHandler) TripHandler {
	return &tripHandler{
		TripUseCase:   TripUseCase,
		CookieHandler: CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool) TripHandler {
	cookieLayer := cd.CreateDelivery(db)
	tripLayer := NewTripHandler(tu.NewTripUseCase(tr.NewTripStorage(db)), cookieLayer)
	return tripLayer
}

func SetUpTripRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	tripHandler := CreateDelivery(db)

	r.GET(cnst.TripURL, tripHandler.Trip)
	r.POST(cnst.TripPostURL, tripHandler.AddTrip)
	r.PATCH(cnst.TripURL, tripHandler.Update)
	r.DELETE(cnst.TripURL, tripHandler.Delete)

	return r
}

func (s *tripHandler) Trip(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, _ := s.CookieHandler.GetUser(hash)
	if !s.TripUseCase.CheckAuthor(foundUser, param) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	code, bytes := s.TripUseCase.GetById(param)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *tripHandler) AddTrip(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		logger.Error("unable to determine user")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	trip := new(domain.Trip)
	if err := json.Unmarshal(ctx.PostBody(), &trip); err != nil {
		logger.Error("error while unmarshalling JSON: ", zap.Error(err))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	id, err := s.TripUseCase.Add(*trip, foundUser)
	if err != nil {
		logger.Error("error while adding trip")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, bytes := s.TripUseCase.GetById(id)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *tripHandler) Update(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, _ := s.CookieHandler.GetUser(hash)
	if !s.TripUseCase.CheckAuthor(foundUser, param) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	trip := new(domain.Trip)
	if err := json.Unmarshal(ctx.PostBody(), &trip); err != nil {
		logger.Error("error while unmarshalling JSON: ", zap.Error(err))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err := s.TripUseCase.Update(param, *trip)
	if err != nil {
		logger.Error("error while updating trip")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, bytes := s.TripUseCase.GetById(param)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *tripHandler) Delete(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, _ := s.CookieHandler.GetUser(hash)
	if !s.TripUseCase.CheckAuthor(foundUser, param) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	err := s.TripUseCase.Delete(param)
	if err != nil {
		logger.Error("error while deleting trip")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}
