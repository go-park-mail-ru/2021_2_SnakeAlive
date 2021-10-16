package delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"snakealive/m/domain"
	"snakealive/m/entities"
	"snakealive/m/repository"
	"time"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

const CookieName = "SnakeAlive"

type SessionServer interface {
	Login(ctx *fasthttp.RequestCtx)
	Registration(ctx *fasthttp.RequestCtx)
	PlacesList(ctx *fasthttp.RequestCtx)
}

type sessionServer struct {
	UserUseCase  domain.UserUseCase
	PlaceUseCase domain.PlaceUseCase
}

func NewSessionServer(UserUseCase domain.UserUseCase, PlaceUseCase domain.PlaceUseCase) SessionServer {
	return &sessionServer{
		UserUseCase:  UserUseCase,
		PlaceUseCase: PlaceUseCase,
	}
}

func (s *sessionServer) Login(ctx *fasthttp.RequestCtx) {
	user := new(domain.User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if !Validate(user) {
		log.Printf("error while validate user:")
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

func (s *sessionServer) Registration(ctx *fasthttp.RequestCtx) {
	user := new(domain.User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if !Validate(user) {
		log.Printf("error while validate user:")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	u, found := s.UserUseCase.Get(user.Email)
	if found {
		log.Printf("User with this email already exists")
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

func Validate(u *domain.User) bool {
	ok, err := regexp.Match(`^\w+[.\w]+@\w+[.\w]+$`, []byte(u.Email))

	if err != nil {
		return false
	}
	if !ok || u.Email == "" {
		return false
	}
	if len(u.Password) < 8 || u.Password == "" || len(u.Password) > 254 || len(u.Email) > 254 {
		return false
	}
	return true
}

func (s *sessionServer) PlacesList(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("name").(string)
	if _, found := s.PlaceUseCase.Get(param); !found {
		log.Printf("country not found")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	response, _ := s.PlaceUseCase.Get(param)
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func Profile(ctx *fasthttp.RequestCtx) {
	if !CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	user := repository.CookieDB[string(ctx.Request.Header.Cookie(CookieName))]
	response := domain.User{Name: user.Name, Surname: user.Surname}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func Logout(ctx *fasthttp.RequestCtx) {
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

	repository.CookieDB[cookie] = user
}

func DeleteCookie(ctx *fasthttp.RequestCtx, cookie string) {
	var c fasthttp.Cookie
	c.SetKey(CookieName)
	c.SetValue("")
	c.SetMaxAge(0)
	c.SetExpire(time.Now().Add(-1))
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)

	delete(repository.CookieDB, cookie)
}

func SetToken(ctx *fasthttp.RequestCtx, hash string) {
	t := entities.Token{Token: hash}
	bytes, err := json.Marshal(t)

	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func CheckCookie(ctx *fasthttp.RequestCtx) bool {
	if _, found := repository.CookieDB[string(ctx.Request.Header.Cookie(CookieName))]; !found {
		return false
	}
	return true
}
