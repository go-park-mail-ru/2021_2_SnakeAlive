package reviewDelivery

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	cd "snakealive/m/internal/cookie/delivery"
// 	ru "snakealive/m/internal/review/usecase"
// 	"snakealive/m/pkg/domain"
// 	service_mocks "snakealive/m/pkg/domain/mocks"
// 	"testing"

// 	"github.com/jackc/pgx/v4/pgxpool"
// 	"github.com/valyala/fasthttp"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// type mockBehavior func(r *service_mocks.MockReviewStorage, user domain.Review)

// type MyTest struct {
// 	name                 string
// 	inputBody            string
// 	inputReview          domain.Review
// 	mockBehavior         mockBehavior
// 	expectedStatusCode   int
// 	expectedResponseBody string
// }

// func SetUpDB() *pgxpool.Pool {
// 	url := "postgres://tripadvisor:12345@localhost:5432/tripadvisor"

// 	dbpool, err := pgxpool.Connect(context.Background(), url)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	return dbpool
// }

// func TestHandler_ReviewsByPlace(t *testing.T) {
// 	tests := []MyTest{
// 		{
// 			name:               "OK",
// 			inputBody:          `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
// 			mockBehavior:       func(r *service_mocks.MockReviewStorage, user domain.Review) {},
// 			expectedStatusCode: fasthttp.StatusBadRequest,
// 		},
// 	}
// 	ctx := &fasthttp.RequestCtx{}

// 	for _, tc := range tests {
// 		c := gomock.NewController(t)
// 		defer c.Finish()
// 		repo := service_mocks.NewMockReviewStorage(c)
// 		tc.mockBehavior(repo, tc.inputReview)

// 		ctx.Request.SetBody(nil)
// 		ctx.Request.AppendBody([]byte(tc.inputBody))
// 		cookieLayer := cd.CreateDelivery(SetUpDB())
// 		userLayer := NewReviewHandler(ru.NewReviewUseCase(repo), cookieLayer)
// 		userLayer.ReviewsByPlace(ctx)

// 		assert.Equal(t, ctx.Response.Header.StatusCode(), tc.expectedStatusCode)
// 	}
// }
