package http

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
	ent "snakealive/m/entities"
	"snakealive/m/session/storage"
	"snakealive/m/validate"
)

type SessionServer interface {
	Login(ctx *fasthttp.RequestCtx)
	Registration(ctx *fasthttp.RequestCtx)
}

type sessionServer struct {
	userStorage storage.UserStorage
}

func NewSessionServer(userStorage storage.UserStorage) SessionServer {
	return &sessionServer{userStorage: userStorage}
}

func (s *sessionServer) Login(ctx *fasthttp.RequestCtx) {
	user := new(ent.User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userData, found := s.userStorage.Get(user.Email)
	if !found {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	if userData.Password != user.Password {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (s *sessionServer) Registration(ctx *fasthttp.RequestCtx) {
	user := new(ent.User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if !validate.Validate(*user) {
		log.Printf("error while validate user:")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if _, found := s.userStorage.Get(user.Email); found {
		log.Printf("User with this email already exists")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	s.userStorage.Add(user.Email, *user)

	ctx.SetStatusCode(fasthttp.StatusOK)
}
