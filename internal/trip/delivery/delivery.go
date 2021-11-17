package tripDelivery

import (
	"encoding/json"
	"snakealive/m/internal/domain"
	"strconv"

	cd "snakealive/m/internal/cookie/delivery"
	tr "snakealive/m/internal/trip/repository"
	tu "snakealive/m/internal/trip/usecase"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type tripHandler struct {
	TripUseCase   domain.TripUseCase
	CookieHandler domain.CookieHandler
}

func NewTripHandler(TripUseCase domain.TripUseCase, CookieHandler domain.CookieHandler) domain.TripHandler {
	return &tripHandler{
		TripUseCase:   TripUseCase,
		CookieHandler: CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool) domain.TripHandler {
	cookieLayer := cd.CreateDelivery(db)
	tripLayer := NewTripHandler(tu.NewTripUseCase(tr.NewTripStorage(db)), cookieLayer)
	return tripLayer
}

func SetUpTripRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	tripHandler := CreateDelivery(db)

	r.GET(cnst.TripURL, logs.AccessLogMiddleware(tripHandler.Trip))
	r.POST(cnst.TripPostURL, logs.AccessLogMiddleware(tripHandler.AddTrip))
	r.PATCH(cnst.TripURL, logs.AccessLogMiddleware(tripHandler.Update))
	r.DELETE(cnst.TripURL, logs.AccessLogMiddleware(tripHandler.Delete))

	return r
}

func (s *tripHandler) Trip(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()

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
}

func (s *tripHandler) AddTrip(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()

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
}

func (s *tripHandler) Update(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()

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
}

func (s *tripHandler) Delete(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()

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

	response := map[string]int{"status": fasthttp.StatusOK}
	bytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("error while marshalling JSON: %s", zap.Error(err))
		return
	}
	ctx.Write(bytes)
}
