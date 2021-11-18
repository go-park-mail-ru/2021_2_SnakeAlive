package cookieDelivery

import (
	"encoding/json"
	cr "snakealive/m/internal/cookie/repository"
	cu "snakealive/m/internal/cookie/usecase"
	"snakealive/m/internal/domain"
	ent "snakealive/m/internal/entities"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type cookieHandler struct {
	CookieUseCase domain.CookieUseCase
}

func NewCookieHandler(CookieUseCase domain.CookieUseCase) domain.CookieHandler {
	return &cookieHandler{
		CookieUseCase: CookieUseCase,
	}
}

func CreateDelivery(db *pgxpool.Pool, l *logs.Logger) domain.CookieHandler {
	cookieLayer := NewCookieHandler(cu.NewCookieUseCase(cr.NewCookieStorage(db), l))
	return cookieLayer
}

func (s *cookieHandler) SetCookieAndToken(ctx *fasthttp.RequestCtx, cookie string, id int) {
	s.SetCookie(ctx, cookie, id)
	setToken(ctx, cookie)
}

func (s *cookieHandler) SetCookie(ctx *fasthttp.RequestCtx, cookie string, id int) {
	var c fasthttp.Cookie
	c.SetKey(cnst.CookieName)
	c.SetValue(cookie)
	c.SetMaxAge(int(time.Hour))
	c.SetHTTPOnly(true)
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)

	s.CookieUseCase.Add(cookie, id)
}

func (s *cookieHandler) DeleteCookie(ctx *fasthttp.RequestCtx, cookie string) {
	var c fasthttp.Cookie
	c.SetKey(cnst.CookieName)
	c.SetValue("")
	c.SetMaxAge(0)
	c.SetExpire(time.Now().Add(-1))
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)

	s.CookieUseCase.Delete(cookie)
}

func (s *cookieHandler) CheckCookie(ctx *fasthttp.RequestCtx) bool {
	cookieHash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	_, err := s.CookieUseCase.Get(cookieHash)
	return err == nil
}

func (s *cookieHandler) GetUser(cookie string) (user domain.User, err error) {
	return s.CookieUseCase.Get(cookie)
}

func setToken(ctx *fasthttp.RequestCtx, hash string) {
	t := ent.Token{
		Token: hash,
	}

	bytes, _ := json.Marshal(t)
	ctx.Write(bytes)
}
