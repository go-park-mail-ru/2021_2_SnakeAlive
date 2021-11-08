package domain

import "github.com/valyala/fasthttp"

type CookieHandler interface {
	SetCookieAndToken(ctx *fasthttp.RequestCtx, cookie string, id int)
	SetCookie(ctx *fasthttp.RequestCtx, cookie string, id int)
	DeleteCookie(ctx *fasthttp.RequestCtx, cookie string)
	CheckCookie(ctx *fasthttp.RequestCtx) bool
	GetUser(cookie string) (user User, err error)
}

type CookieStorage interface {
	Add(key string, userId int) error
	Get(value string) (user User, err error)
	Delete(value string) error
}

type CookieUseCase interface {
	Add(key string, userId int) error
	Get(value string) (user User, err error)
	Delete(value string) error
}
