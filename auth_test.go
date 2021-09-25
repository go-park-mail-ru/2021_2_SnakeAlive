package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

type test struct {
	input string
	want  int
}

func TestLogin(t *testing.T) {
	tests := []test{
		{(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "password": "password"}`),
			fasthttp.StatusOK},
		{(`{"name": "name2", "surname": "surname2", "email": "alextest@mail.ru", "password": "password"}`),
			fasthttp.StatusNotFound},
		{(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "password": "xxxxxxxx"}`),
			fasthttp.StatusBadRequest},
		{(`xxxxx`), fasthttp.StatusBadRequest},
	}
	var c fasthttp.RequestCtx
	ctx := (*fasthttp.RequestCtx)(&c)

	for _, tc := range tests {
		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.input))
		login(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}
}

func TestRegister(t *testing.T) {
	tests := []test{
		{(`{"name": "name2", "surname": "surname2", "email": "asdf@mail.ru", "password": "pass"}`),
			fasthttp.StatusOK},
		{(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "password": "pass"}`),
			fasthttp.StatusBadRequest},
		{(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "xxxxx": "xxxxxxxx"}`),
			fasthttp.StatusBadRequest},
		{(`{"name": "name2", "surname": "surname2", "xxxx": "alex@mail.ru", "password": "xxxxxxxx"}`),
			fasthttp.StatusBadRequest},
		{(`xxxxx`), fasthttp.StatusBadRequest},
	}
	var c fasthttp.RequestCtx
	ctx := (*fasthttp.RequestCtx)(&c)

	for _, tc := range tests {
		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.input))
		registration(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}
}

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestPlacesList(t *testing.T) {

	tests := []test{
		{"Russia", fasthttp.StatusOK},
		{"Germany", fasthttp.StatusNotFound},
	}
	var c fasthttp.RequestCtx
	ctx := (*fasthttp.RequestCtx)(&c)

	ctx.Request.AppendBody([]byte(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "password": "password"}`))
	login(ctx)
	assert.Equal(t, fasthttp.StatusOK, ctx.Response.Header.StatusCode())

	placesList(ctx)
	assert.Equal(t, fasthttp.StatusUnauthorized, ctx.Response.Header.StatusCode())

	for _, tc := range tests {
		ctx.Request.Header.SetCookie("SnakeAlive", "3259306991")
		ctx.SetUserValue("name", tc.input)
		placesList(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}

}
