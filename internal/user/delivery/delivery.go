package userDelivery

import (
	"encoding/json"
	"fmt"
	"log"
	ent "snakealive/m/internal/entities"
	cnst "snakealive/m/pkg/constants"
	chttp "snakealive/m/pkg/customhttp"
	"snakealive/m/pkg/domain"

	ur "snakealive/m/internal/user/repository"
	uu "snakealive/m/internal/user/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/fasthttp/router"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

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

func NewUserHandler(UserUseCase domain.UserUseCase) UserHandler {
	return &userHandler{
		UserUseCase: UserUseCase,
	}
}

func CreateDelivery(db *pgxpool.Pool) UserHandler {
	userLayer := NewUserHandler(uu.NewUserUseCase(ur.NewUserStorage(db)))
	return userLayer
}

func SetUpUserRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	userHandler := CreateDelivery(db)
	r.POST(cnst.LoginURL, userHandler.Login)
	r.POST(cnst.RegisterURL, userHandler.Registration)
	r.GET(cnst.ProfileURL, userHandler.GetProfile)
	r.PATCH(cnst.ProfileURL, userHandler.UpdateProfile)
	r.DELETE(cnst.ProfileURL, userHandler.DeleteProfile)
	r.DELETE(cnst.LogoutURL, userHandler.Logout)
	return r
}

func (s *userHandler) Login(ctx *fasthttp.RequestCtx) {
	user := new(domain.User)
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Println("!found")
		return
	}

	valid := s.UserUseCase.Validate(user)
	if !valid {
		log.Printf("error while validating user")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, _ := s.UserUseCase.Login(user)
	ctx.SetStatusCode(code)

	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	chttp.SetCookieAndToken(ctx, с, user.Id)
}

func (s *userHandler) Registration(ctx *fasthttp.RequestCtx) {
	user := new(domain.User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, _ := s.UserUseCase.Registration(user)
	ctx.SetStatusCode(code)

	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	newUser, _ := s.UserUseCase.GetByEmail(user.Email)
	chttp.SetCookieAndToken(ctx, с, newUser.Id)
}

func (s *userHandler) GetProfile(ctx *fasthttp.RequestCtx) {
	if !chttp.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	id := ent.CookieDB[string(ctx.Request.Header.Cookie(cnst.CookieName))]
	foundUser, err := s.UserUseCase.GetById(id)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	response := map[string]string{"name": foundUser.Name, "surname": foundUser.Surname}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func (s *userHandler) UpdateProfile(ctx *fasthttp.RequestCtx) {
	if !chttp.CheckCookie(ctx) {
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

	id := ent.CookieDB[string(ctx.Request.Header.Cookie(cnst.CookieName))]
	foundUser, err := s.UserUseCase.GetById(id)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	if err = s.UserUseCase.Update(foundUser.Id, *updatedUser); err != nil {
		log.Printf("user with this email already exists")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	response := map[string]string{"name": updatedUser.Name, "surname": updatedUser.Surname, "email": updatedUser.Email}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func (s *userHandler) DeleteProfile(ctx *fasthttp.RequestCtx) {
	if !chttp.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	id := ent.CookieDB[string(ctx.Request.Header.Cookie(cnst.CookieName))]
	foundUser, err := s.UserUseCase.GetById(id)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	s.UserUseCase.Delete(foundUser.Id)
	chttp.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
}

func (s *userHandler) Logout(ctx *fasthttp.RequestCtx) {
	if !chttp.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	chttp.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
}
