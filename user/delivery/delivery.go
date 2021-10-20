package userDelivery

import (
	"encoding/json"
	"fmt"
	"log"
	"snakealive/m/domain"
	ent "snakealive/m/entities"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/fasthttp/router"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

const CookieName = "SnakeAlive"

type UserHandler interface {
	Login(ctx *fasthttp.RequestCtx)
	Registration(ctx *fasthttp.RequestCtx)
	Logout(ctx *fasthttp.RequestCtx)
	GetProfile(ctx *fasthttp.RequestCtx)
	UpdateProfile(ctx *fasthttp.RequestCtx)
	DeleteProfile(ctx *fasthttp.RequestCtx)
}

type userHandler struct {
	UserUseCase domain.UserUseCase
}

func NewUserHandler(UserUseCase domain.UserUseCase, r *router.Router) {
	userHandler := userHandler{UserUseCase: UserUseCase}

	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Registration)
	r.DELETE("/logout", userHandler.Logout)
	r.GET("/profile", userHandler.GetProfile)
	r.PATCH("/profile", userHandler.UpdateProfile)
	r.DELETE("/profile", userHandler.DeleteProfile)
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

func (s *userHandler) GetProfile(ctx *fasthttp.RequestCtx) {
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

func (s *userHandler) UpdateProfile(ctx *fasthttp.RequestCtx) {
	if !CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	updatedUser := new(domain.User)
	if err := json.Unmarshal(ctx.PostBody(), &updatedUser); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	_, err := govalidator.ValidateStruct(updatedUser)
	if err != nil {
		log.Printf("error while validating user")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	currentUser := ent.CookieDB[string(ctx.Request.Header.Cookie(CookieName))]

	if !s.UserUseCase.Update(currentUser, *updatedUser) {
		log.Printf("user with this email already exists")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	bytes, err := json.Marshal(updatedUser)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func (s *userHandler) DeleteProfile(ctx *fasthttp.RequestCtx) {
	if !CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	user := ent.CookieDB[string(ctx.Request.Header.Cookie(CookieName))]
	s.UserUseCase.Delete(user.Email)
	DeleteCookie(ctx, string(ctx.Request.Header.Cookie(CookieName)))
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
