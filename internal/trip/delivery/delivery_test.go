package tripDelivery

import (
	"context"
	"fmt"
	"os"
	cu "snakealive/m/internal/cookie/usecase"
	tu "snakealive/m/internal/trip/usecase"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/domain"
	service_mocks "snakealive/m/pkg/domain/mocks"
	"testing"

	cd "snakealive/m/internal/cookie/delivery"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

//type mockBehavior func(r *service_mocks.MockTripStorage, trip domain.Trip, user domain.User)

type MyTest struct {
	name                 string
	inputBody            string
	inputTrip            domain.Trip
	inputUser            domain.User
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

func TestHandler_AddTrip(t *testing.T) {
	user := domain.User{
		Id:       1,
		Name:     "Александра",
		Surname:  "Волкова",
		Email:    "testtesttests@mail.ru",
		Password: "1234567890",
	}
	tests := []MyTest{
		/*{
			name:      "Unathorised",
			inputBody: `{"title": "Best trip ever", "description": "Wow wow", "days": [[ {"id": 1}, {"id": 2} ] ]}`,
			inputTrip: domain.Trip{
				Title:       "Best trip ever",
				Description: "Wow wow",
				Days:        [][]domain.Place{{domain.Place{Id: 1}, domain.Place{Id: 2}}},
			},
			inputUser:          domain.User{},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},*/
		{
			name:      "Trip added",
			inputBody: `{"title": "Best trip ever", "description": "Wow wow", "days": [[ {"id": 1}, {"id": 2} ] ]}`,
			inputTrip: domain.Trip{
				Id:          0,
				Title:       "Best trip ever",
				Description: "Wow wow",
				Days:        [][]domain.Place{{domain.Place{Id: 1}, domain.Place{Id: 2}}},
			},
			inputUser:          user,
			expectedStatusCode: fasthttp.StatusOK,
		},
	}

	mockAdd := func(r *service_mocks.MockTripStorage, trip domain.Trip, user domain.User) {
		r.EXPECT().Add(trip, gomock.Any()).Return(1, nil).AnyTimes()
	}

	mockGetUser := func(r *service_mocks.MockCookieStorage, cookie string) {
		r.EXPECT().Get(gomock.Any()).Return(user, nil).AnyTimes()
	}

	mockGetById := func(r *service_mocks.MockTripStorage, trip domain.Trip, id int) {
		r.EXPECT().GetById(gomock.Any()).Return(trip, nil).AnyTimes()
	}

	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		cookieRepo := service_mocks.NewMockCookieStorage(c)
		cookie := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email)))

		tripRepo := service_mocks.NewMockTripStorage(c)
		mockAdd(tripRepo, tc.inputTrip, tc.inputUser)
		mockGetUser(cookieRepo, cookie)
		mockGetById(tripRepo, tc.inputTrip, tc.inputTrip.Id)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))

		if tc.expectedStatusCode != fasthttp.StatusUnauthorized {
			ctx.Request.Header.SetCookie(cnst.CookieName, cookie)
		}

		cookieLayer := cd.NewCookieHandler(cu.NewCookieUseCase(cookieRepo))
		tripLayer := NewTripHandler(tu.NewTripUseCase(tripRepo), cookieLayer)
		tripLayer.AddTrip(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}
