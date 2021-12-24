package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"snakealive/m/internal/gateway/user/usecase"
	"snakealive/m/internal/models"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"

	"github.com/valyala/fasthttp"
)

type UserDelivery interface {
	Login(ctx *fasthttp.RequestCtx)
	Logout(ctx *fasthttp.RequestCtx)
	Register(ctx *fasthttp.RequestCtx)
	GetUserInfo(ctx *fasthttp.RequestCtx)

	GetProfile(ctx *fasthttp.RequestCtx)
	UpdateProfile(ctx *fasthttp.RequestCtx)
}

type userDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.UserUsecase
}

func (u *userDelivery) Login(ctx *fasthttp.RequestCtx) {
	var request = &models.LoginUserRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := u.manager.Login(ctx, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := json.Marshal(response)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)

	var c fasthttp.Cookie
	c.SetKey(cnst.CookieName)
	c.SetValue(response.Cookie)
	c.SetMaxAge(int(time.Hour))
	c.SetHTTPOnly(true)
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)
}

func (u *userDelivery) Logout(ctx *fasthttp.RequestCtx) {
	err := u.manager.Logout(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	ctx.SetStatusCode(http.StatusOK)
}

func (u *userDelivery) Register(ctx *fasthttp.RequestCtx) {
	var request = &models.RegisterUserRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := u.manager.Register(ctx, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := json.Marshal(response)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)

	var c fasthttp.Cookie
	c.SetKey(cnst.CookieName)
	c.SetValue(response.Cookie)
	c.SetMaxAge(int(time.Hour))
	c.SetHTTPOnly(true)
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)
}

func (u *userDelivery) GetProfile(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)
	response, err := u.manager.Profile(ctx, userID)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := json.Marshal(response)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *userDelivery) UpdateProfile(ctx *fasthttp.RequestCtx) {
	var request = &models.UpdateProfileRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	userID := ctx.UserValue(cnst.UserIDContextKey).(int)
	response, err := u.manager.UpdateProfile(ctx, userID, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := json.Marshal(response)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *userDelivery) GetUserInfo(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	responce, err := u.manager.GetUserInfo(ctx, param)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := json.Marshal(responce)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func NewUserDelivery(
	errorAdapter error_adapter.HttpAdapter,
	manager usecase.UserUsecase,
) UserDelivery {
	return &userDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
