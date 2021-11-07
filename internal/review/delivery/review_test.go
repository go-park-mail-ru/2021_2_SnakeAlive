package reviewDelivery

import (
	"context"
	"errors"
	"fmt"
	"os"
	cd "snakealive/m/internal/cookie/delivery"
	cu "snakealive/m/internal/cookie/usecase"
	logs "snakealive/m/internal/logger"
	ru "snakealive/m/internal/review/usecase"
	ud "snakealive/m/internal/user/delivery"
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

type mockReviewBehavior func(r *service_mocks.MockReviewStorage, review domain.Review)
type mockUserBehavior func(r *service_mocks.MockUserStorage, user domain.User)

type MyTest struct {
	name                 string
	urlArg               int
	inputBody            string
	inputReview          domain.Review
	outputUser           domain.User
	mockReviewBehavior   mockReviewBehavior
	mockUserBehavior     mockUserBehavior
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
			mockReviewBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {
				r.EXPECT().GetListByPlace(review.PlaceId, 0, 10)
			},
			mockUserBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetById(gomock.Any()).Return(user, nil).AnyTimes()
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
			outputUser: domain.User{
				Id:       1,
				Name:     "Александра",
				Surname:  "Волкова",
				Email:    "testtesttests@mail.ru",
				Password: "1234567890",
			},
			mockReviewBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {
				r.EXPECT().GetListByPlace(review.PlaceId, 0, 10).Return(domain.ReviewsNoPlace{}, errors.New("err"))
			},
			mockUserBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetById(gomock.Any()).Return(user, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusNotFound,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		logs.BuildLogger()

		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockReviewBehavior(repo, tc.inputReview)

		userRepo := service_mocks.NewMockUserStorage(c)
		tc.mockUserBehavior(userRepo, tc.outputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := ud.CreateDelivery(SetUpDB())
		reviewLayer := NewReviewHandler(ru.NewReviewUseCase(repo, userRepo), cookieLayer, userLayer)

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))
		reviewLayer.ReviewsByPlace(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_AddReviewToPlace(t *testing.T) {
	user := domain.User{
		Id:       1,
		Name:     "Александра",
		Surname:  "Волкова",
		Email:    "testtesttests@mail.ru",
		Password: "1234567890",
	}
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"title": "title", "text": "text", "rating": 10, "place_id": 1}`,
			urlArg:    1,
			inputReview: domain.Review{
				Title:   "title",
				Text:    "text",
				Rating:  10,
				PlaceId: 1,
			},
			mockReviewBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {
				r.EXPECT().Add(review, review.PlaceId)
			},
			mockUserBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetById(gomock.Any()).Return(user, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
	}
	ctx := &fasthttp.RequestCtx{}
	mockGetUser := func(r *service_mocks.MockCookieStorage, cookie string, user domain.User) {
		r.EXPECT().Get(gomock.Any()).Return(user, nil).AnyTimes()
	}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		logs.BuildLogger()

		cookieRepo := service_mocks.NewMockCookieStorage(c)
		cookie := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
		mockGetUser(cookieRepo, cookie, user)

		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockReviewBehavior(repo, tc.inputReview)

		userRepo := service_mocks.NewMockUserStorage(c)
		tc.mockUserBehavior(userRepo, tc.outputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := ud.CreateDelivery(SetUpDB())
		reviewLayer := NewReviewHandler(ru.NewReviewUseCase(repo, userRepo), cookieLayer, userLayer)

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))

		ctx.Request.Header.SetCookie(cnst.CookieName, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte("alex@mail.ru"))))
		reviewLayer.AddReviewToPlace(ctx)

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
			mockReviewBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {},
			mockUserBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetById(gomock.Any()).Return(user, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		logs.BuildLogger()

		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockReviewBehavior(repo, tc.inputReview)

		userRepo := service_mocks.NewMockUserStorage(c)
		tc.mockUserBehavior(userRepo, tc.outputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := ud.CreateDelivery(SetUpDB())
		reviewLayer := NewReviewHandler(ru.NewReviewUseCase(repo, userRepo), cookieLayer, userLayer)

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))

		reviewLayer.AddReviewToPlace(ctx)

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
			mockReviewBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {},
			mockUserBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetById(gomock.Any()).Return(user, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		logs.BuildLogger()

		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockReviewBehavior(repo, tc.inputReview)

		userRepo := service_mocks.NewMockUserStorage(c)
		tc.mockUserBehavior(userRepo, tc.outputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := ud.CreateDelivery(SetUpDB())
		reviewLayer := NewReviewHandler(ru.NewReviewUseCase(repo, userRepo), cookieLayer, userLayer)

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))

		reviewLayer.DelReview(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_DelReview2(t *testing.T) {
	user := domain.User{
		Id:       1,
		Name:     "Александра",
		Surname:  "Волкова",
		Email:    "testtesttests@mail.ru",
		Password: "1234567890",
	}
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"title": "title", "text": "text", "rating": 10, "place_id": 1}`,
			urlArg:    1,
			inputReview: domain.Review{
				Title:   "title",
				Text:    "text",
				Rating:  10,
				PlaceId: 1,
			},
			mockReviewBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {

				r.EXPECT().GetReviewAuthor(1).Return(1)
				r.EXPECT().Get(1)
				r.EXPECT().Delete(1)
			},
			mockUserBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetById(gomock.Any()).Return(user, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
		{
			name:      "not author",
			inputBody: `{"title": "title", "text": "text", "rating": 10,"place_id": 1}`,
			urlArg:    1,
			inputReview: domain.Review{
				Title:   "title",
				Text:    "text",
				Rating:  10,
				PlaceId: 1,
			},
			mockReviewBehavior: func(r *service_mocks.MockReviewStorage, review domain.Review) {

				r.EXPECT().GetReviewAuthor(1).Return(0)
			},
			mockUserBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetById(gomock.Any()).Return(user, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	mockGetUser := func(r *service_mocks.MockCookieStorage, cookie string, user domain.User) {
		r.EXPECT().Get(cookie).Return(user, nil).AnyTimes()
	}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()
		logs.BuildLogger()

		repo := service_mocks.NewMockReviewStorage(c)
		tc.mockReviewBehavior(repo, tc.inputReview)

		userRepo := service_mocks.NewMockUserStorage(c)
		tc.mockUserBehavior(userRepo, tc.outputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))

		ctx.SetUserValue("id", fmt.Sprint(tc.urlArg))

		cookieRepo := service_mocks.NewMockCookieStorage(c)
		cookie := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
		mockGetUser(cookieRepo, cookie, user)

		cookieLayer := cd.NewCookieHandler(cu.NewCookieUseCase(cookieRepo))
		userLayer := ud.CreateDelivery(SetUpDB())
		reviewLayer := NewReviewHandler(ru.NewReviewUseCase(repo, userRepo), cookieLayer, userLayer)

		ctx.Request.Header.SetCookie(cnst.CookieName, cookie)
		reviewLayer.DelReview(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}
