package userDelivery

import (
	"encoding/json"
	"fmt"
	"snakealive/m/internal/domain"
	cnst "snakealive/m/pkg/constants"

	cd "snakealive/m/internal/cookie/delivery"
	ur "snakealive/m/internal/user/repository"
	uu "snakealive/m/internal/user/usecase"
	logs "snakealive/m/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type userHandler struct {
	UserUseCase   domain.UserUseCase
	CookieHandler domain.CookieHandler
}

func NewUserHandler(UserUseCase domain.UserUseCase, CookieHandler domain.CookieHandler) domain.UserHandler {
	return &userHandler{
		UserUseCase:   UserUseCase,
		CookieHandler: CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool, l *logs.Logger) domain.UserHandler {
	cookieLayer := cd.CreateDelivery(db, l)
	userLayer := NewUserHandler(uu.NewUserUseCase(ur.NewUserStorage(db), l), cookieLayer)
	return userLayer
}

func SetUpUserRouter(db *pgxpool.Pool, r *router.Router, l *logs.Logger) *router.Router {
	userHandler := CreateDelivery(db, l)
	r.POST(cnst.LoginURL, logs.AccessLogMiddleware(l, userHandler.Login))
	r.POST(cnst.RegisterURL, logs.AccessLogMiddleware(l, userHandler.Registration))
	r.GET(cnst.ProfileURL, logs.AccessLogMiddleware(l, userHandler.GetProfile))
	r.PATCH(cnst.ProfileURL, logs.AccessLogMiddleware(l, userHandler.UpdateProfile))
	r.DELETE(cnst.ProfileURL, logs.AccessLogMiddleware(l, userHandler.DeleteProfile))
	r.DELETE(cnst.LogoutURL, logs.AccessLogMiddleware(l, userHandler.Logout))
	r.POST(cnst.UploadURL, logs.AccessLogMiddleware(l, userHandler.UploadAvatar))
	return r
}

func (s *userHandler) Bind(user *domain.User, ctx *fasthttp.RequestCtx) {
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	valid := s.UserUseCase.Validate(user)
	if !valid {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
}

func (s *userHandler) Login(ctx *fasthttp.RequestCtx) {
	user := new(domain.User)
	s.Bind(user, ctx)

	code, err := s.UserUseCase.Login(user)
	ctx.SetStatusCode(code)
	if err != nil {
		return
	}

	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	foundUser, _ := s.UserUseCase.GetByEmail(user.Email)
	s.CookieHandler.SetCookieAndToken(ctx, с, foundUser.Id)
}

func (s *userHandler) Registration(ctx *fasthttp.RequestCtx) {

	user := new(domain.User)

	s.Bind(user, ctx)

	code, err := s.UserUseCase.Registration(user)
	ctx.SetStatusCode(code)
	if err != nil {
		return
	}

	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	newUser, _ := s.UserUseCase.GetByEmail(user.Email)
	s.CookieHandler.SetCookieAndToken(ctx, с, newUser.Id)
}

func (s *userHandler) GetProfile(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	code, bytes := s.UserUseCase.GetProfile(hash, foundUser)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}

func (s *userHandler) UpdateProfile(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	updatedUser := new(domain.User)
	s.Bind(updatedUser, ctx)

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	code, bytes := s.UserUseCase.UpdateProfile(updatedUser, foundUser, hash)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}

func (s *userHandler) DeleteProfile(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)

	s.UserUseCase.DeleteProfile(hash, foundUser)
	s.CookieHandler.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))

	response := map[string]int{"status": fasthttp.StatusOK}
	bytes, err := json.Marshal(response)
	if err != nil {
		return
	}
	ctx.Write(bytes)
}

func (s *userHandler) DeleteProfileByEmail(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	s.UserUseCase.DeleteUserByEmail(foundUser)
	s.CookieHandler.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
}

func (s *userHandler) Logout(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	s.CookieHandler.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	response := map[string]int{"status": fasthttp.StatusOK}
	bytes, err := json.Marshal(response)
	if err != nil {
		return
	}
	ctx.Write(bytes)
}

func (s *userHandler) UploadAvatar(ctx *fasthttp.RequestCtx) {
	if !s.CookieHandler.CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	avatar := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(foundUser.Email))) + cnst.Format

	formFile, err := ctx.FormFile(cnst.FormFileKey)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = fasthttp.SaveMultipartFile(formFile, cnst.StaticPath+cnst.AvatarDirPath+avatar)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	avatarURL := cnst.StaticServerURL + cnst.AvatarDirPath + avatar
	code, bytes := s.UserUseCase.AddAvatar(foundUser, avatarURL)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
}
