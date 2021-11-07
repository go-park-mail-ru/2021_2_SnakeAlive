package reviewDelivery

import (
	"encoding/json"
	"fmt"
	cd "snakealive/m/internal/cookie/delivery"
	logs "snakealive/m/internal/logger"
	ud "snakealive/m/internal/user/delivery"
	ur "snakealive/m/internal/user/repository"
	"snakealive/m/pkg/domain"
	"strconv"

	rr "snakealive/m/internal/review/repository"
	ru "snakealive/m/internal/review/usecase"
	cnst "snakealive/m/pkg/constants"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type ReviewHandler interface {
	ReviewsByPlace(ctx *fasthttp.RequestCtx)
	AddReviewToPlace(ctx *fasthttp.RequestCtx)
	DelReview(ctx *fasthttp.RequestCtx)
}

type reviewHandler struct {
	ReviewUseCase domain.ReviewUseCase
	CookieHandler cd.CookieHandler
	UserHandler   ud.UserHandler
}

func NewReviewHandler(ReviewUseCase domain.ReviewUseCase, CookieHandler cd.CookieHandler, UserHandler ud.UserHandler) ReviewHandler {
	return &reviewHandler{
		ReviewUseCase: ReviewUseCase,
		CookieHandler: CookieHandler,
		UserHandler:   UserHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool) ReviewHandler {
	cookieLayer := cd.CreateDelivery(db)
	userLayer := ud.CreateDelivery(db)
	reviewLayer := NewReviewHandler(ru.NewReviewUseCase(rr.NewReviewStorage(db), ur.NewUserStorage(db)), cookieLayer, userLayer)
	return reviewLayer
}

func SetUpReviewRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	reviewHandler := CreateDelivery(db)
	r.POST(cnst.ReviewAddURL, reviewHandler.AddReviewToPlace)
	r.GET(cnst.ReviewURL, reviewHandler.ReviewsByPlace)
	r.DELETE(cnst.ReviewURL, reviewHandler.DelReview)
	return r
}

func (s *reviewHandler) ReviewsByPlace(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

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
	fmt.Println("limit skip = ", limit, skip)
	code, bytes := s.ReviewUseCase.GetReviewsListByPlaceId(id, user, limit, skip)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *reviewHandler) AddReviewToPlace(ctx *fasthttp.RequestCtx) {
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
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *reviewHandler) DelReview(ctx *fasthttp.RequestCtx) {
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
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}
