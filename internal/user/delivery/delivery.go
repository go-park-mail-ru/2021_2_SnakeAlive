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
	"go.uber.org/zap"
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

func CreateDelivery(db *pgxpool.Pool) domain.UserHandler {
	cookieLayer := cd.CreateDelivery(db)
	userLayer := NewUserHandler(uu.NewUserUseCase(ur.NewUserStorage(db)), cookieLayer)
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
	r.POST(cnst.UploadURL, userHandler.UploadAvatar)
	return r
}

func (s *userHandler) Bind(user *domain.User, ctx *fasthttp.RequestCtx, logger *zap.Logger) {
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		logger.Error("error while unmarshalling JSON: ", zap.Error(err))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	valid := s.UserUseCase.Validate(user)
	if !valid {
		logger.Error("error while validating user")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
}

func (s *userHandler) Login(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	user := new(domain.User)
	s.Bind(user, ctx, logger)

	code, err := s.UserUseCase.Login(user)
	ctx.SetStatusCode(code)
	if err != nil {
		logger.Error("error while logging user in")
		return
	}

	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	foundUser, _ := s.UserUseCase.GetByEmail(user.Email)
	s.CookieHandler.SetCookieAndToken(ctx, с, foundUser.Id)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *userHandler) Registration(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	user := new(domain.User)

	s.Bind(user, ctx, logger)

	code, err := s.UserUseCase.Registration(user)
	ctx.SetStatusCode(code)
	if err != nil {
		logger.Error("error while registering user")
		return
	}

	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	newUser, _ := s.UserUseCase.GetByEmail(user.Email)
	s.CookieHandler.SetCookieAndToken(ctx, с, newUser.Id)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *userHandler) GetProfile(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))

	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		logger.Error("unable to determine user")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	code, bytes := s.UserUseCase.GetProfile(hash, foundUser)

	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *userHandler) UpdateProfile(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	updatedUser := new(domain.User)
	s.Bind(updatedUser, ctx, logger)

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		logger.Error("uunable to determine user")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	code, bytes := s.UserUseCase.UpdateProfile(updatedUser, foundUser, hash)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *userHandler) DeleteProfile(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		logger.Error("unable to determine user")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	s.UserUseCase.DeleteProfile(hash, foundUser)
	s.CookieHandler.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *userHandler) DeleteProfileByEmail(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		logger.Error("uable to determine user")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	s.UserUseCase.DeleteUserByEmail(foundUser)
	s.CookieHandler.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *userHandler) Logout(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	s.CookieHandler.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}

func (s *userHandler) UploadAvatar(ctx *fasthttp.RequestCtx) {
	logger := logs.GetLogger()
	logger.Info(string(ctx.Path()),
		zap.String("method", string(ctx.Method())),
		zap.String("remote_addr", string(ctx.RemoteAddr().String())),
		zap.String("url", string(ctx.Path())),
	)

	if !s.CookieHandler.CheckCookie(ctx) {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	hash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	foundUser, err := s.CookieHandler.GetUser(hash)
	if err != nil {
		logger.Error("user is unauthorized")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	avatar := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(foundUser.Email))) + cnst.Format

	formFile, err := ctx.FormFile(cnst.FormFileKey)
	if err != nil {
		logger.Error("no formfile found")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = fasthttp.SaveMultipartFile(formFile, cnst.StaticPath+"/"+avatar)
	if err != nil {
		logger.Error("unable to save file")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, bytes := s.UserUseCase.AddAvatar(foundUser, avatar)
	ctx.SetStatusCode(code)
	ctx.Write(bytes)
	logger.Info(string(ctx.Path()),
		zap.Int("status", ctx.Response.StatusCode()),
	)
}
