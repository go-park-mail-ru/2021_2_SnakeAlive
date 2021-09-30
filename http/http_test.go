package http

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
		{(`{"name": "name2", "surname": "surname2", "email": "alexmail.ru", "password": "xxxxxxxx"}`),
			fasthttp.StatusBadRequest},
		{(`xxxxx`), fasthttp.StatusBadRequest},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.input))
		Login(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}
}

func TestRegister(t *testing.T) {
	tests := []test{
		{(`{"name": "name2", "surname": "surname2", "email": "asdf@mail.ru", "password": "12345678"}`),
			fasthttp.StatusOK},
		{(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "password": "12345678"}`),
			fasthttp.StatusBadRequest},
		{(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "xxxxx": "xxxxxxxx"}`),
			fasthttp.StatusBadRequest},
		{(`{"name": "name2", "surname": "surname2", "xxxx": "alex@mail.ru", "password": "xxxxxxxx"}`),
			fasthttp.StatusBadRequest},
		{(`{"name": "name2", "surname": "surname2", "xxxx": "alex@mail.ru", "password": "1234567"}`),
			fasthttp.StatusBadRequest},
		{(`{"name": "name2", "surname": "surname2", "xxxx": "alexmail.ru", "password": "1234567"}`),
			fasthttp.StatusBadRequest},
		{(`xxxxx`), fasthttp.StatusBadRequest},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.input))
		Registration(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}
}

func TestPlacesList(t *testing.T) {

	tests := []test{
		{"Russia", fasthttp.StatusOK},
		{"Nicaragua", fasthttp.StatusOK},
		{"Germany", fasthttp.StatusNotFound},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		ctx.Request.Header.SetCookie("SnakeAlive", "3259306991")
		ctx.SetUserValue("name", tc.input)
		PlacesList(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}

}
