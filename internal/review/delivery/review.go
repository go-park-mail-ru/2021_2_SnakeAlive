package reviewDelivery

import (
	"encoding/json"
	"log"
	cd "snakealive/m/internal/cookie/delivery"
	"snakealive/m/pkg/domain"
	"strconv"

	rr "snakealive/m/internal/review/repository"
	ru "snakealive/m/internal/review/usecase"
	cnst "snakealive/m/pkg/constants"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type ReviewHandler interface {
	ReviewsByPlace(ctx *fasthttp.RequestCtx)
	AddReviewToPlace(ctx *fasthttp.RequestCtx)
	DelReview(ctx *fasthttp.RequestCtx)
}

type reviewHandler struct {
	ReviewUseCase domain.ReviewUseCase
	CookieHandler cd.CookieHandler
}

func NewReviewHandler(ReviewUseCase domain.ReviewUseCase, CookieHandler cd.CookieHandler) ReviewHandler {
	return &reviewHandler{
		ReviewUseCase: ReviewUseCase,
		CookieHandler: CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool) ReviewHandler {
	cookieLayer := cd.CreateDelivery(db)
	reviewLayer := NewReviewHandler(ru.NewReviewUseCase(rr.NewReviewStorage(db)), cookieLayer)
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
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	code, bytes := s.ReviewUseCase.GetReviewsListByPlaceId(id)
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
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, err := s.ReviewUseCase.Add(*newReview, foundUser)
	ctx.SetStatusCode(code)
	if err != nil {
		log.Printf("error while registering user in")
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
		log.Printf("No such review")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	s.ReviewUseCase.Delete(id)
}
