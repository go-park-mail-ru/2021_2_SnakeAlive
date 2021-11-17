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
	"go.uber.org/zap"
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

func CreateDelivery(db *pgxpool.Pool) domain.ReviewHandler {
	cookieLayer := cd.CreateDelivery(db)
	userLayer := ud.CreateDelivery(db)
	reviewLayer := NewReviewHandler(ru.NewReviewUseCase(rr.NewReviewStorage(db), ur.NewUserStorage(db)), cookieLayer, userLayer)
	return reviewLayer
}

func SetUpReviewRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	reviewHandler := CreateDelivery(db)
	r.POST(cnst.ReviewAddURL, logs.AccessLogMiddleware(reviewHandler.AddReviewToPlace))
	r.GET(cnst.ReviewURL, logs.AccessLogMiddleware(reviewHandler.ReviewsByPlace))
	r.DELETE(cnst.ReviewURL, logs.AccessLogMiddleware(reviewHandler.DelReview))
	return r
}

func (s *reviewHandler) ReviewsByPlace(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()

	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	cookieHash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	var user domain.User
	user, err := s.CookieHandler.GetUser(cookieHash)
	if err != nil {
		logger.Error("unable to determine user")
		user = domain.User{}
	}

	skip, err := strconv.Atoi(string(ctx.QueryArgs().Peek("skip")))
	if err != nil {
		logger.Error("unable to get query arg skip")
		skip = cnst.DefaultSkip
	}

	limit, err := strconv.Atoi(string(ctx.QueryArgs().Peek("limit")))
	if err != nil {
		logger.Error("unable to get query arg limit")
		limit = 0
	}

	code, bytes := s.ReviewUseCase.GetReviewsListByPlaceId(id, user, limit, skip)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}

func (s *reviewHandler) AddReviewToPlace(ctx *fasthttp.RequestCtx) {
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

	newReview := new(domain.Review)

	if err := json.Unmarshal(ctx.PostBody(), &newReview); err != nil {
		logger.Error("error while unmarshalling JSON: %s", zap.Error(err))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, bytes, err := s.ReviewUseCase.Add(*newReview, foundUser)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	if err != nil {
		logger.Error("error while registering user in")
		return
	}
}

func (s *reviewHandler) DelReview(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil || !s.ReviewUseCase.CheckAuthor(foundUser, id) {
		logger.Error("user doesn't have permission for this action")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	_, err = s.ReviewUseCase.Get(id)
	if err != nil {
		logger.Error("review not found")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	s.ReviewUseCase.Delete(id)

	response := map[string]int{"status": fasthttp.StatusOK}
	bytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("error while marshalling JSON: %s", zap.Error(err))
		return
	}
	ctx.Write(bytes)
}
