package userDelivery

import (
	"context"
	"errors"
	"fmt"
	"os"
	cd "snakealive/m/internal/cookie/delivery"
	uu "snakealive/m/internal/user/usecase"
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

type mockBehavior func(r *service_mocks.MockUserStorage, user domain.User)

type MyTest struct {
	name                 string
	inputBody            string
	inputUser            domain.User
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

func TestHandler_Login(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Name:     "Name",
				Surname:  "surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetByEmail(user.Email).Return(user, nil)
				r.EXPECT().GetByEmail(user.Email).Return(user, nil)
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
		{
			name:      "Validate error",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "123"}`,
			inputUser: domain.User{},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
			},
			expectedStatusCode: fasthttp.StatusBadRequest,
		},
		{
			name:      "Wrong password",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Id:       0,
				Name:     "Name",
				Surname:  "surname",
				Email:    "alex@mail.ru",
				Password: "12345667",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetByEmail(user.Email).Return(user, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusBadRequest,
		},
		{
			name:      "No such user",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex123231@mail.ru", "password": "password"}`,
			inputUser: domain.User{},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetByEmail("alex123231@mail.ru").Return(user, errors.New("Error")).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusNotFound,
		},
		{
			name:               "Json unmarshalling error",
			inputBody:          `---`,
			inputUser:          domain.User{},
			mockBehavior:       func(r *service_mocks.MockUserStorage, user domain.User) {},
			expectedStatusCode: fasthttp.StatusBadRequest,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)
		userLayer.Login(ctx)

		assert.Equal(t, ctx.Response.Header.StatusCode(), tc.expectedStatusCode)
	}
}

func TestHandler_Register(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "testing@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Name:     "Name",
				Surname:  "Surname",
				Email:    "testing@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetByEmail(user.Email).Return(domain.User{}, errors.New("Error")).AnyTimes()
				r.EXPECT().Add(user)
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
		{
			name:               "Json unmarshal error",
			inputBody:          `---`,
			inputUser:          domain.User{},
			mockBehavior:       func(r *service_mocks.MockUserStorage, user domain.User) {},
			expectedStatusCode: fasthttp.StatusBadRequest,
		},
		{
			name:               "Validate error",
			inputBody:          `{"----": "Name", "surname": "Surname", "email": "testing@mail.ru", "password": "password"}`,
			inputUser:          domain.User{},
			mockBehavior:       func(r *service_mocks.MockUserStorage, user domain.User) {},
			expectedStatusCode: fasthttp.StatusBadRequest,
		},
		{
			name:      "User registered",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetByEmail(user.Email).Return(domain.User{}, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusBadRequest,
		},
		{
			name:               "Json unmarshalling error",
			inputBody:          `---`,
			inputUser:          domain.User{},
			mockBehavior:       func(r *service_mocks.MockUserStorage, user domain.User) {},
			expectedStatusCode: fasthttp.StatusBadRequest,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)
		userLayer.Registration(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_Logout(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior:       func(r *service_mocks.MockUserStorage, user domain.User) {},
			expectedStatusCode: fasthttp.StatusBadRequest,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)
		userLayer.Logout(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_Logout2(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Id:       1,
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior:       func(r *service_mocks.MockUserStorage, user domain.User) {},
			expectedStatusCode: fasthttp.StatusOK,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)

		ctx.Request.Header.SetCookie(cnst.CookieName, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))))
		cookieLayer.SetCookieAndToken(ctx, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))), tc.inputUser.Id)

		userLayer.Logout(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}
func TestHandler_GetProfile(t *testing.T) {
	tests := []MyTest{
		{
			name:      "Unauthorized",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetByEmail(user.Email).Return(domain.User{}, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)

		userLayer.GetProfile(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}

}

func TestHandler_GetProfile2(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Id:       1,
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().GetByEmail(user.Email).Return(domain.User{}, nil).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)

		ctx.Request.Header.SetCookie(cnst.CookieName, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))))
		cookieLayer.SetCookieAndToken(ctx, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))), tc.inputUser.Id)

		userLayer.GetProfile(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}

}

func TestHandler_UpdateProfile(t *testing.T) {
	tests := []MyTest{
		{
			name:      "Unauthorized",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Id:       1,
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior:       func(r *service_mocks.MockUserStorage, user domain.User) {},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)

		//ctx.Request.Header.SetCookie(cnst.CookieName, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))))
		userLayer.UpdateProfile(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_UpdateProfile2(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Id:       0,
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().Update(1, user).AnyTimes()
				r.EXPECT().GetByEmail(user.Email).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)

		ctx.Request.Header.SetCookie(cnst.CookieName, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))))
		cookieLayer.SetCookieAndToken(ctx, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))), tc.inputUser.Id)

		userLayer.UpdateProfile(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_DeleteProfile(t *testing.T) {
	tests := []MyTest{
		{
			name:      "Unauthorized",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Id:       1,
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().Delete(user.Id).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)

		userLayer.DeleteProfile(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_DeleteProfile2(t *testing.T) {
	tests := []MyTest{
		{
			name:      "OK",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Id:       1,
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().Delete(user.Id).AnyTimes()
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)

		ctx.Request.Header.SetCookie(cnst.CookieName, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))))
		cookieLayer.SetCookieAndToken(ctx, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))), tc.inputUser.Id)

		userLayer.DeleteProfile(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_DeleteProfileByEmail2(t *testing.T) {
	tests := []MyTest{
		{
			name:      "Unauthorized",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Id:       1,
				Name:     "name",
				Surname:  "surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {
				r.EXPECT().DeleteByEmail(user)
			},
			expectedStatusCode: fasthttp.StatusOK,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)

		ctx.Request.Header.SetCookie(cnst.CookieName, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))))
		cookieLayer.SetCookieAndToken(ctx, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))), tc.inputUser.Id)

		userLayer.DeleteProfileByEmail(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}

func TestHandler_DeleteProfileByEmail(t *testing.T) {
	tests := []MyTest{
		{
			name:      "Unauthorized",
			inputBody: `{"name": "Name", "surname": "Surname", "email": "alex@mail.ru", "password": "password"}`,
			inputUser: domain.User{
				Id:       1,
				Name:     "Name",
				Surname:  "Surname",
				Email:    "alex@mail.ru",
				Password: "password",
			},
			mockBehavior: func(r *service_mocks.MockUserStorage, user domain.User) {

			},
			expectedStatusCode: fasthttp.StatusUnauthorized,
		},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		c := gomock.NewController(t)
		defer c.Finish()

		repo := service_mocks.NewMockUserStorage(c)
		tc.mockBehavior(repo, tc.inputUser)

		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.inputBody))
		cookieLayer := cd.CreateDelivery(SetUpDB())
		userLayer := NewUserHandler(uu.NewUserUseCase(repo), cookieLayer)

		// ctx.Request.Header.SetCookie(cnst.CookieName, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))))
		// cookieLayer.SetCookieAndToken(ctx, fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(tc.inputUser.Email))), tc.inputUser.Id)

		userLayer.DeleteProfileByEmail(ctx)

		assert.Equal(t, tc.expectedStatusCode, ctx.Response.Header.StatusCode())
	}
}
