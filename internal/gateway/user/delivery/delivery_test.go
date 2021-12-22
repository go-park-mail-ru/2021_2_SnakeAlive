package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"snakealive/m/internal/gateway/user/usecase"
	"snakealive/m/internal/models"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
	auth_service "snakealive/m/pkg/services/auth"
	mock_auth_service "snakealive/m/pkg/services/auth/mock"
)

type Test struct {
	Prepare func(cli *mock_auth_service.MockAuthServiceClient)
	Run     func(d UserDelivery, t *testing.T)
}

const (
	defaultUserID = 1
	cookie        = "cookie"

	defaultUserName        = "defaultUserName"
	defaultUserSurname     = "defaultUserSurname"
	defaultUserEmail       = "defaultUserEmail"
	defaultUserImage       = "defaultUserImage"
	defaultUserDescription = "defaultUserDescription"
	pass                   = "pass"
)

var (
	someError = errors.New("error")

	loginRequest = &models.LoginUserRequest{
		Email:    defaultUserEmail,
		Password: pass,
	}
	registerRequest = &models.RegisterUserRequest{
		Email:    defaultUserEmail,
		Password: pass,
		Name:     defaultUserName,
		Surname:  defaultUserSurname,
	}
	updateProfileRequest = &models.UpdateProfileRequest{
		Name:        defaultUserName,
		Surname:     defaultUserSurname,
		Avatar:      defaultUserImage,
		Email:       defaultUserEmail,
		Description: defaultUserDescription,
		Password:    pass,
	}

	defaultErrorMsg = "msg"
	defaultError    = error_adapter.HttpError{
		MSG:  defaultErrorMsg,
		Code: http.StatusTeapot,
	}

	tests = []Test{
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().GetUser(gomock.Any(), &auth_service.GetUserRequest{Id: int64(defaultUserID)}).Return(nil, someError)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				d.GetProfile(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().GetUser(gomock.Any(), &auth_service.GetUserRequest{Id: int64(defaultUserID)}).Return(&auth_service.GetUserResponse{
					Name:        defaultUserName,
					Surname:     defaultUserSurname,
					Email:       defaultUserEmail,
					Image:       defaultUserImage,
					Description: defaultUserDescription,
				}, nil)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				d.GetProfile(ctx)

				var profile = models.Profile{
					Id:          defaultUserID,
					Name:        defaultUserName,
					Surname:     defaultUserSurname,
					Avatar:      defaultUserImage,
					Email:       defaultUserEmail,
					Description: defaultUserDescription,
				}
				var resp models.Profile

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
				getResp(ctx, t, &resp)
				assert.Equal(t, resp, profile)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()

				d.Login(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
				assert.Equal(t, string(ctx.Response.Body()), cnst.WrongRequestBody)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().LoginUser(gomock.Any(), &auth_service.LoginRequest{
					Email:    defaultUserEmail,
					Password: pass,
				}).Return(nil, someError)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx, t, loginRequest)
				d.Login(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().LoginUser(gomock.Any(), &auth_service.LoginRequest{
					Email:    defaultUserEmail,
					Password: pass,
				}).Return(&auth_service.LoginResponse{
					Cookie: "123",
					Token:  "??",
				}, nil)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx, t, loginRequest)
				d.Login(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().LogoutUser(gomock.Any(), &auth_service.Session{
					Token:  "??",
					Cookie: cookie,
				}).Return(nil, someError)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				d.Logout(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().LogoutUser(gomock.Any(), &auth_service.Session{
					Token:  "??",
					Cookie: cookie,
				}).Return(nil, nil)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				d.Logout(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()

				d.Register(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
				assert.Equal(t, string(ctx.Response.Body()), cnst.WrongRequestBody)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().RegisterUser(gomock.Any(), &auth_service.RegisterRequest{
					Email:    defaultUserEmail,
					Password: pass,
					Name:     defaultUserName,
					Surname:  defaultUserSurname,
				}).Return(nil, someError)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx, t, registerRequest)
				d.Register(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().RegisterUser(gomock.Any(), &auth_service.RegisterRequest{
					Email:    defaultUserEmail,
					Password: pass,
					Name:     defaultUserName,
					Surname:  defaultUserSurname,
				}).Return(&auth_service.LoginResponse{
					Cookie: "123",
					Token:  "??",
				}, nil)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx, t, registerRequest)
				d.Register(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()

				d.UpdateProfile(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
				assert.Equal(t, string(ctx.Response.Body()), cnst.WrongRequestBody)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().UpdateUser(gomock.Any(), &auth_service.UpdateUserRequest{
					Id:          defaultUserID,
					Name:        defaultUserName,
					Surname:     defaultUserSurname,
					Email:       defaultUserEmail,
					Description: defaultUserDescription,
					Password:    pass,
					Image:       defaultUserImage,
				}).Return(nil, someError)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx, t, updateProfileRequest)
				d.UpdateProfile(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().UpdateUser(gomock.Any(), &auth_service.UpdateUserRequest{
					Id:          defaultUserID,
					Name:        defaultUserName,
					Surname:     defaultUserSurname,
					Email:       defaultUserEmail,
					Description: defaultUserDescription,
					Password:    pass,
					Image:       defaultUserImage,
				}).Return(&auth_service.GetUserResponse{
					Name:        defaultUserName,
					Surname:     defaultUserSurname,
					Email:       defaultUserEmail,
					Image:       defaultUserImage,
					Description: defaultUserDescription,
				}, nil)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx, t, updateProfileRequest)
				d.UpdateProfile(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().GetUser(gomock.Any(), &auth_service.GetUserRequest{Id: int64(defaultUserID)}).Return(nil, someError)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", defaultUserID)
				d.GetProfile(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_auth_service.MockAuthServiceClient) {
				cli.EXPECT().GetUser(gomock.Any(), &auth_service.GetUserRequest{Id: int64(defaultUserID)}).Return(&auth_service.GetUserResponse{
					Name:        defaultUserName,
					Surname:     defaultUserSurname,
					Email:       defaultUserEmail,
					Image:       defaultUserImage,
					Description: defaultUserDescription,
				}, nil)
			},
			Run: func(d UserDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", defaultUserID)
				d.GetProfile(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
	}
)

func TestDelivery(t *testing.T) {
	for i := range tests {
		d, cli := prepare(t)
		tests[i].Prepare(cli)
		tests[i].Run(d, t)
	}
}

func prepare(t *testing.T) (d UserDelivery, cli *mock_auth_service.MockAuthServiceClient) {
	ctrl := gomock.NewController(t)
	cli = mock_auth_service.NewMockAuthServiceClient(ctrl)
	d = NewUserDelivery(
		error_adapter.NewErrorToHttpAdapter(map[error]error_adapter.HttpError{}, defaultError),
		usecase.NewUserUsecase(cli),
	)

	return
}

func getCtx() *fasthttp.RequestCtx {
	ctx := &fasthttp.RequestCtx{}
	setUserDefaultCtx(ctx)
	return ctx
}

func setUserDefaultCtx(ctx *fasthttp.RequestCtx) {
	ctx.SetUserValue(cnst.UserIDContextKey, defaultUserID)
	ctx.Request.Header.SetCookie(cnst.CookieName, cookie)
}

func setBody(ctx *fasthttp.RequestCtx, t *testing.T, val interface{}) {
	b, err := json.Marshal(val)
	assert.NoError(t, err)

	ctx.Request.SetBody(b)
}

func getResp(ctx *fasthttp.RequestCtx, t *testing.T, val interface{}) {
	assert.NoError(t, json.Unmarshal(ctx.Response.Body(), val))
}
