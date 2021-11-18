package reviewDelivery

import (
	"encoding/json"
	cd "snakealive/m/internal/cookie/delivery"
	"snakealive/m/internal/domain"
	ud "snakealive/m/internal/user/delivery"
	ur "snakealive/m/internal/user/repository"
	logs "snakealive/m/pkg/logger"
	"strconv"

	rr "snakealive/m/internal/review/repository"
	ru "snakealive/m/internal/review/usecase"
	cnst "snakealive/m/pkg/constants"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type reviewHandler struct {
	ReviewUseCase domain.ReviewUseCase
	CookieHandler domain.CookieHandler
	UserHandler   domain.UserHandler
}

func NewReviewHandler(ReviewUseCase domain.ReviewUseCase, CookieHandler domain.CookieHandler, UserHandler domain.UserHandler) domain.ReviewHandler {
	return &reviewHandler{
		ReviewUseCase: ReviewUseCase,
		CookieHandler: CookieHandler,
		UserHandler:   UserHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool, l *logs.Logger) domain.ReviewHandler {
	cookieLayer := cd.CreateDelivery(db, l)
	userLayer := ud.CreateDelivery(db, l)
	reviewLayer := NewReviewHandler(ru.NewReviewUseCase(rr.NewReviewStorage(db), ur.NewUserStorage(db), l), cookieLayer, userLayer)
	return reviewLayer
}

func SetUpReviewRouter(db *pgxpool.Pool, r *router.Router, l *logs.Logger) *router.Router {
	reviewHandler := CreateDelivery(db, l)
	r.POST(cnst.ReviewAddURL, logs.AccessLogMiddleware(l, reviewHandler.AddReviewToPlace))
	r.GET(cnst.ReviewURL, logs.AccessLogMiddleware(l, reviewHandler.ReviewsByPlace))
	r.DELETE(cnst.ReviewURL, logs.AccessLogMiddleware(l, reviewHandler.DelReview))
	return r
}

func (s *reviewHandler) ReviewsByPlace(ctx *fasthttp.RequestCtx) {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	cookieHash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	var user domain.User
	user, err := s.CookieHandler.GetUser(cookieHash)
	if err != nil {
		user = domain.User{}
	}

	skip, err := strconv.Atoi(string(ctx.QueryArgs().Peek("skip")))
	if err != nil {
		skip = cnst.DefaultSkip
	}

	limit, err := strconv.Atoi(string(ctx.QueryArgs().Peek("limit")))
	if err != nil {
		limit = 0
	}

	code, bytes := s.ReviewUseCase.GetReviewsListByPlaceId(id, user, limit, skip)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}

func (s *reviewHandler) AddReviewToPlace(ctx *fasthttp.RequestCtx) {
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

	newReview := new(domain.Review)

	if err := json.Unmarshal(ctx.PostBody(), &newReview); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, bytes, err := s.ReviewUseCase.Add(*newReview, foundUser)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	if err != nil {
		return
	}
}

func (s *reviewHandler) DelReview(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil || !s.ReviewUseCase.CheckAuthor(foundUser, id) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	_, err = s.ReviewUseCase.Get(id)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	s.ReviewUseCase.Delete(id)

	response := map[string]int{"status": fasthttp.StatusOK}
	bytes, err := json.Marshal(response)
	if err != nil {
		return
	}
	ctx.Write(bytes)
}
