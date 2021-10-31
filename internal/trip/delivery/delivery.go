package tripDelivery

import (
	"encoding/json"
	"log"
	"snakealive/m/pkg/domain"
	"strconv"

	cd "snakealive/m/internal/cookie/delivery"
	tr "snakealive/m/internal/trip/repository"
	tu "snakealive/m/internal/trip/usecase"
	cnst "snakealive/m/pkg/constants"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
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

func SetUpPlaceRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	tripHandler := CreateDelivery(db)

	r.GET(cnst.TripURL, tripHandler.Trip)
	r.POST(cnst.TripPostURL, tripHandler.AddTrip)
	r.PATCH(cnst.TripURL, tripHandler.Update)
	r.DELETE(cnst.TripURL, tripHandler.Delete)

	return r
}

func (s *tripHandler) Trip(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, _ := s.CookieHandler.GetUser(hash)
	if !s.TripUseCase.CheckAuthor(foundUser, param) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	code, bytes := s.TripUseCase.GetById(param)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}

func (s *tripHandler) AddTrip(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	trip := new(domain.Trip)
	if err := json.Unmarshal(ctx.PostBody(), &trip); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	id, err := s.TripUseCase.Add(*trip, foundUser)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, bytes := s.TripUseCase.GetById(id)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}

func (s *tripHandler) Update(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, _ := s.CookieHandler.GetUser(hash)
	if !s.TripUseCase.CheckAuthor(foundUser, param) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	trip := new(domain.Trip)
	if err := json.Unmarshal(ctx.PostBody(), &trip); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err := s.TripUseCase.Update(param, *trip)
	if err != nil {
		log.Printf("error updating: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, bytes := s.TripUseCase.GetById(param)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}

func (s *tripHandler) Delete(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, _ := s.CookieHandler.GetUser(hash)
	if !s.TripUseCase.CheckAuthor(foundUser, param) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	err := s.TripUseCase.Delete(param)
	if err != nil {
		log.Printf("error updating: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
}
