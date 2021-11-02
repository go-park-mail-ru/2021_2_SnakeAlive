package cookieDelivery

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	cr "snakealive/m/internal/cookie/repository"
	cu "snakealive/m/internal/cookie/usecase"
	logs "snakealive/m/internal/logger"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/domain"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type CookieHandler interface {
	SetCookieAndToken(ctx *fasthttp.RequestCtx, cookie string, user domain.User)
	setCookie(ctx *fasthttp.RequestCtx, cookie string, id int)
	CheckCookieAndToken(ctx *fasthttp.RequestCtx, token string, user domain.User) bool
	DeleteCookie(ctx *fasthttp.RequestCtx, cookie string)
	CheckCookie(ctx *fasthttp.RequestCtx) bool
	GetUser(cookie string) (user domain.User, err error)
}

type cookieHandler struct {
	CookieUseCase domain.CookieUseCase
}

func NewCookieHandler(CookieUseCase domain.CookieUseCase) CookieHandler {
	return &cookieHandler{
		CookieUseCase: CookieUseCase,
	}
}

func CreateDelivery(db *pgxpool.Pool) CookieHandler {
	cookieLayer := NewCookieHandler(cu.NewCookieUseCase(cr.NewCookieStorage(db)))
	return cookieLayer
}

func (s *cookieHandler) SetCookieAndToken(ctx *fasthttp.RequestCtx, cookie string, user domain.User) {
	logger := logs.GetLogger()
	s.setCookie(ctx, cookie, user.Id)
	tokens, err := NewHMACHashToken(cnst.TokenName)
	if err != nil {
		logger.Error("unable to create token")
	}
	tokens.Create(ctx, user, time.Now().Add(24*time.Hour).Unix())
}

func (s *cookieHandler) CheckCookieAndToken(ctx *fasthttp.RequestCtx, token string, user domain.User) bool {
	logger := logs.GetLogger()

	tokens, err := NewHMACHashToken(cnst.TokenName)
	if err != nil {
		logger.Error("unable to create token")
	}
	rigthToken, err := tokens.Check(ctx, user, token)
	if err != nil {
		logger.Error("token error")
	}
	return s.CheckCookie(ctx) && rigthToken
}

func (s *cookieHandler) setCookie(ctx *fasthttp.RequestCtx, cookie string, id int) {
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
	logger := logs.GetLogger()
	cookieHash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	_, err := s.CookieUseCase.Get(cookieHash)
	if err != nil {
		logger.Error("unable to find cookie")
	}
	return err == nil
}

func (s *cookieHandler) GetUser(cookie string) (user domain.User, err error) {
	return s.CookieUseCase.Get(cookie)
}

type HashToken struct {
	Secret []byte
}

func NewHMACHashToken(secret string) (*HashToken, error) {
	return &HashToken{Secret: []byte(secret)}, nil
}

func (tk *HashToken) Create(ctx *fasthttp.RequestCtx, user domain.User, tokenExpTime int64) {
	logger := logs.GetLogger()

	h := hmac.New(sha256.New, []byte(tk.Secret))
	data := fmt.Sprintf("%d:%s:%d", user.Id, user.Email, tokenExpTime)
	h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(tokenExpTime, 10)

	bytes, err := json.Marshal(token)
	if err != nil {
		logger.Error("error while marshalling JSON: ", zap.Error(err))
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func (tk *HashToken) Check(ctx *fasthttp.RequestCtx, user domain.User, inputToken string) (bool, error) {
	tokenData := strings.Split(inputToken, ":")
	if len(tokenData) != 2 {
		return false, fmt.Errorf("bad token data")
	}

	tokenExp, err := strconv.ParseInt(tokenData[1], 10, 64)
	if err != nil {
		return false, fmt.Errorf("bad token time")
	}

	if tokenExp < time.Now().Unix() {
		return false, fmt.Errorf("token expired")
	}

	h := hmac.New(sha256.New, []byte(tk.Secret))
	data := fmt.Sprintf("%d:%s:%d", user.Id, user.Email, tokenExp)
	h.Write([]byte(data))
	expectedMAC := h.Sum(nil)
	messageMAC, err := hex.DecodeString(tokenData[0])
	if err != nil {
		return false, fmt.Errorf("cand hex decode token")
	}

	return hmac.Equal(messageMAC, expectedMAC), nil
}
