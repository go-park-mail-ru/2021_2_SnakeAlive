package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

type User struct {
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty"`
}

var auth = map[string]User{
	"alex@mail.ru": {Name: "alex", Surname: "surname", Email: "alex@mail.ru", Password: "pass"},
}

type Token struct {
	Token string `json:"token,omitempty"`
}

func handlerDecorator(handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
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
	if !bytes.Equal(ctx.Method(), []byte(fasthttp.MethodPost)) {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	user := new(User)
	//err := validate.Validate(User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if _, found := auth[user.Email]; !found {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	password := auth[user.Email].Password

	if password != user.Password {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	t := Token{"token"}
	bytes, err := json.Marshal(t)

	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func registration(ctx *fasthttp.RequestCtx) {

	if !bytes.Equal(ctx.Method(), []byte(fasthttp.MethodPost)) {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	user := new(User)
	//err := validate.Validate(User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if _, found := auth[user.Email]; found {
		log.Printf("User with this email already exists")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	auth[user.Email] = *user

	ctx.SetStatusCode(fasthttp.StatusOK)
	t := Token{"token"}
	bytes, err := json.Marshal(t)

	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func SetCookie(ctx *fasthttp.RequestCtx) {
	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Println(auth)
	switch string(ctx.Path()) {
	case "/login":
		login(ctx)
	case "/register":
		registration(ctx)
	default:
		fmt.Println("no rout")
	}
}

func main() {
	fmt.Println("starting server at :8080")
	if err := fasthttp.ListenAndServe(":8080", handlerDecorator(requestHandler)); err != nil {
		fmt.Println("failed to start server:", err)
		return
	}
}
