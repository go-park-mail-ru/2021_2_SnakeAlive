package userDelivery

import (
	"encoding/json"
	"fmt"
	"log"
	ent "snakealive/m/internal/entities"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/domain"
	"time"

	ur "snakealive/m/internal/user/repository"
	uu "snakealive/m/internal/user/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/fasthttp/router"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

const CookieName = "SnakeAlive"

type UserHandler interface {
	Login(ctx *fasthttp.RequestCtx)
	Registration(ctx *fasthttp.RequestCtx)
	Profile(ctx *fasthttp.RequestCtx)
	Logout(ctx *fasthttp.RequestCtx)
}

type userHandler struct {
	UserUseCase domain.UserUseCase
}

func NewUserHandler(UserUseCase domain.UserUseCase) UserHandler {
	return &userHandler{
		UserUseCase: UserUseCase,
	}
}

func CreateDelivery() UserHandler {
	userLayer := NewUserHandler(uu.NewUserUseCase(ur.NewUserStorage()))
	return userLayer
}

func SetUpUserRouter(r *router.Router) *router.Router {
	userHandler := CreateDelivery()
	r.POST(cnst.LOGIN, userHandler.Login)
	r.POST(cnst.REGISTER, userHandler.Registration)
	r.GET(cnst.PROFILE, userHandler.Profile)
	r.DELETE(cnst.LOGOUT, userHandler.Logout)
	return r
}

func (s *userHandler) Login(ctx *fasthttp.RequestCtx) {
	user := new(domain.User)
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	valid := validateLogin(user)
	if !valid {
		log.Printf("error while validating user")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	u, found := s.UserUseCase.Get(user.Email)

	u = *user
	if !found {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	if u.Password != user.Password {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	SetCookie(ctx, с, u)
	SetToken(ctx, с)
}

func (s *userHandler) Registration(ctx *fasthttp.RequestCtx) {
	user := new(domain.User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		log.Printf("error while validating user")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	u, found := s.UserUseCase.Get(user.Email)
	if found {
		log.Printf("user with this email already exists")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	u = *user
	s.UserUseCase.Add(u)
	ctx.SetStatusCode(fasthttp.StatusOK)
	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	SetCookie(ctx, с, u)
	SetToken(ctx, с)
}

func (s *userHandler) Profile(ctx *fasthttp.RequestCtx) {
	if !CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	user := ent.CookieDB[string(ctx.Request.Header.Cookie(CookieName))]
	response := domain.User{Name: user.Name, Surname: user.Surname}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func (s *userHandler) Logout(ctx *fasthttp.RequestCtx) {
	if !CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	DeleteCookie(ctx, string(ctx.Request.Header.Cookie(CookieName)))
}

func SetCookie(ctx *fasthttp.RequestCtx, cookie string, user domain.User) {
	var c fasthttp.Cookie
	c.SetKey(CookieName)
	c.SetValue(cookie)
	c.SetMaxAge(36000)
	c.SetHTTPOnly(true)
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)

	ent.CookieDB[cookie] = user
}

func DeleteCookie(ctx *fasthttp.RequestCtx, cookie string) {
	var c fasthttp.Cookie
	c.SetKey(CookieName)
	c.SetValue("")
	c.SetMaxAge(0)
	c.SetExpire(time.Now().Add(-1))
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)

	delete(ent.CookieDB, cookie)
}

func SetToken(ctx *fasthttp.RequestCtx, hash string) {
	t := ent.Token{Token: hash}
	bytes, err := json.Marshal(t)

	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func CheckCookie(ctx *fasthttp.RequestCtx) bool {
	if _, found := ent.CookieDB[string(ctx.Request.Header.Cookie(CookieName))]; !found {
		return false
	}
	return true
}

func validateLogin(user *domain.User) bool {
	if !govalidator.IsEmail(user.Email) || !govalidator.StringLength(user.Password, "8", "254") ||
		!govalidator.MaxStringLength(user.Email, "254") {
		return false
	}
	return true
}
