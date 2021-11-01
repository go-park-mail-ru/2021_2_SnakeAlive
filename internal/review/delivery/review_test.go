package reviewDelivery

import (
	"context"
	"errors"
	"fmt"
	"os"
	cd "snakealive/m/internal/cookie/delivery"
	ru "snakealive/m/internal/review/usecase"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/domain"
	service_mocks "snakealive/m/pkg/domain/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mockBehavior func(r *service_mocks.MockReviewStorage, user domain.Review)

type MyTest struct {
	name                 string
	urlArg               int
	inputBody            string
	inputReview          domain.Review
	mockBehavior         mockBehavior
	expectedStatusCode   int
	expectedResponseBody string
}

func SetUpDB() *pgxpool.Pool {
	url := "postgres://tripadvisor:12345@localhost:5432/tripadvisor"

	dbpool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}

func TestHandler_ReviewsByPlace(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"title": "title", "text": "text", "rating": 10, "user_id": 1, "place_id": 1}`,
			urlArg:    1,
			inputReview: domain.Review{
				Id:      1,
				Title:   "title",
				Text:    "text",
				Rating:  10,
				UserId:  1,
				PlaceId: 1,
			},
			mockBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {
				r.EXPECT().GetListByPlace(review.PlaceId)
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
		{
			name:      "StatusNotFound reviews",
			inputBody: `{"title": "title", "text": "text", "rating": 10, "user_id": 1, "place_id": 1000}`,
			urlArg:    1000,
			inputReview: domain.Review{
				Id:      1,
				Title:   "title",
				Text:    "text",
				Rating:  10,
				UserId:  1,
				PlaceId: 1000,
			},
			mockBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {
				r.EXPECT().GetListByPlace(review.PlaceId).Return(domain.Reviews{}, errors.New("err"))
			},
			expectedStatusCode: fasthttp.StatusNotFound,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockBehavior(repo, tc.inputReview)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewReviewHandler(ru.NewReviewUseCase(repo), cookieLayer)

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))
		userLayer.ReviewsByPlace(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_AddReviewToPlace(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"title": "title", "text": "text", "rating": 10, "user_id": 1, "place_id": 1}`,
			urlArg:    1,
			inputReview: domain.Review{
				Title:   "title",
				Text:    "text",
				Rating:  10,
				UserId:  1,
				PlaceId: 1,
			},
			mockBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {
				r.EXPECT().Add(review, review.PlaceId)
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
		{
			name:      "OK",
			inputBody: `-`,
			urlArg:    1,
			inputReview: domain.Review{
				Title:   "title",
				Text:    "text",
				Rating:  10,
				UserId:  1,
				PlaceId: 1,
			},
			mockBehavior:       func(r *service_mocks.MockReviewStorage, review domain.Review) {},
			expectedStatusCode: fasthttp.StatusBadRequest,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockBehavior(repo, tc.inputReview)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewReviewHandler(ru.NewReviewUseCase(repo), cookieLayer)

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))

		ctx.Request.Header.SetCookie(cnst.CookieName, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte("alex@mail.ru"))))
		cookieLayer.SetCookieAndToken(ctx, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte("alex@mail.ru"))), 1)

		userLayer.AddReviewToPlace(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_AddReviewToPlace2(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"title": "title", "text": "text", "rating": 10, "user_id": 1, "place_id": 1}`,
			urlArg:    1,
			inputReview: domain.Review{
				Title:   "title",
				Text:    "text",
				Rating:  10,
				UserId:  1,
				PlaceId: 1,
			},
			mockBehavior:       func(r *service_mocks.MockReviewStorage, review domain.Review) {},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockBehavior(repo, tc.inputReview)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewReviewHandler(ru.NewReviewUseCase(repo), cookieLayer)

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))

		userLayer.AddReviewToPlace(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_DelReview(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"title": "title", "text": "text", "rating": 10, "user_id": 1, "place_id": 1}`,
			urlArg:    1,
			inputReview: domain.Review{
				Title:   "title",
				Text:    "text",
				Rating:  10,
				UserId:  1,
				PlaceId: 1,
			},
			mockBehavior:       func(r *service_mocks.MockReviewStorage, review domain.Review) {},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockBehavior(repo, tc.inputReview)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewReviewHandler(ru.NewReviewUseCase(repo), cookieLayer)

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))

		userLayer.DelReview(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_DelReview2(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"title": "title", "text": "text", "rating": 10, "user_id": 1, "place_id": 1}`,
			urlArg:    1,
			inputReview: domain.Review{
				Title:   "title",
				Text:    "text",
				Rating:  10,
				UserId:  1,
				PlaceId: 1,
			},
			mockBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {

				r.EXPECT().GetReviewAuthor(1).Return(1)
				r.EXPECT().Get(1)
				r.EXPECT().Delete(1)
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
		{
			name:      "not author",
			inputBody: `{"title": "title", "text": "text", "rating": 10, "user_id": 1, "place_id": 1}`,
			urlArg:    1,
			inputReview: domain.Review{
				Title:   "title",
				Text:    "text",
				Rating:  10,
				UserId:  1,
				PlaceId: 1,
			},
			mockBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {

				r.EXPECT().GetReviewAuthor(1).Return(0)
			},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockBehavior(repo, tc.inputReview)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewReviewHandler(ru.NewReviewUseCase(repo), cookieLayer)

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))

		ctx.Request.Header.SetCookie(cnst.CookieName, "d37ed98a-c110-37ed-97d5-6f103efc8d7a")
		cookieLayer.SetCookieAndToken(ctx, "d37ed98a-c110-37ed-97d5-6f103efc8d7a", 1)

		userLayer.DelReview(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}
