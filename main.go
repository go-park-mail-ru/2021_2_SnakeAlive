package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"strconv"

	"snakealive/m/auth"
	"snakealive/m/validate"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var authdb = map[string]auth.User{
	"alex@mail.ru": {Name: "alex", Surname: "surname", Email: "alex@mail.ru", Password: "pass"},
}

var CookieDB = map[string]auth.User{}

func Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 10)
}

func corsMiddleware(handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Content-Type", "application/json; charset=utf8")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Request-ID")
		ctx.Response.Header.Set("Access-Control-Expose-Headers", "Authorization")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Max-Age", "3600")

		if bytes.Equal(ctx.Method(), []byte(fasthttp.MethodOptions)) {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		handler(ctx)
	}
}

func login(ctx *fasthttp.RequestCtx) {
	user := new(auth.User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if _, found := authdb[user.Email]; !found {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	password := authdb[user.Email].Password

	if password != user.Password {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	SetCookie(ctx, Hash(user.Email), authdb[user.Email])
}

func registration(ctx *fasthttp.RequestCtx) {
	user := new(auth.User)

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

	if _, found := authdb[user.Email]; found {
		log.Printf("User with this email already exists")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	authdb[user.Email] = *user

	ctx.SetStatusCode(fasthttp.StatusOK)
	SetCookie(ctx, Hash(user.Email), authdb[user.Email])
}

func SetCookie(ctx *fasthttp.RequestCtx, cookie string, user auth.User) {
	var c fasthttp.Cookie
	c.SetKey("SnakeAlive")
	c.SetValue(cookie)
	ctx.Response.Header.SetCookie(&c)

	CookieDB[cookie] = user
}

func Router() *router.Router {
	r := router.New()
	r.POST("/login", login)
	r.POST("/register", registration)
	return r
}

func main() {
	fmt.Println("starting server at :8080")

	r := Router()

	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
		fmt.Println("failed to start server:", err)
		return
	}
}
