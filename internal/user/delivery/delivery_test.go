package userDelivery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"snakealive/m/pkg/domain"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

type test struct {
	input string
	want  int
}

func SetUpDB() *pgxpool.Pool {
	url := "postgres://tripadvisor:12345@localhost:5432/tripadvisor"

	dbpool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer dbpool.Close()
	return dbpool
}

func Test_Login_Success_HttpResponseCode(t *testing.T) {
	tests := []test{
		{`{"email": "alex@mail.ru", "password": "password"}`,
			fasthttp.StatusOK},
		{`{"email": "nikita@mail.ru", "password": "frontend123"}`,
			fasthttp.StatusOK},
		{`{"email": "ksenia@mail.ru", "password": "12345678"}`,
			fasthttp.StatusOK},
		{`{"email": "andrew@mail.ru", "password": "000111000"}`,
			fasthttp.StatusOK},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.input))
		userHandler := CreateDelivery(SetUpDB())
		userHandler.Login(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}
}

func Test_Login_False_HttpResponseCode(t *testing.T) {
	tests := []test{
		{`{"email": "alex@mail.ru", "password": "123123123abcabc"}`,
			fasthttp.StatusBadRequest},
		{`{"email": "------@mail.ru", "password": "--123---"}`,
			fasthttp.StatusNotFound},
		{`{"email": "alex@mail.ru", "password": "-----"}`,
			fasthttp.StatusBadRequest},
		{`{"email": "------", "password": "frontend123"}`,
			fasthttp.StatusBadRequest},
		{`{"-----": "ksenia@mail.ru", "password": "12345678"}`,
			fasthttp.StatusBadRequest},
		{`{--------}`,
			fasthttp.StatusBadRequest},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.input))
		userHandler := CreateDelivery(SetUpDB())
		userHandler.Login(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}
}

func Test_Register_Fail_HttpResponseCode(t *testing.T) {
	tests := []test{
		{`{"name": "123", "surname": "surname", "email": "alex@mail.ru", "password": "password"}`,
			fasthttp.StatusBadRequest},
		{`{"name": "name", "surname": "surname", "email": "alex@mail.ru", "password": "password"}`,
			fasthttp.StatusBadRequest},
		{`{"name": "name2", "surname": "123", "email": "alextest@mail.ru", "password": "password"}`,
			fasthttp.StatusBadRequest},
		{`{"name": "name2", "surname": "surname2", "email": "ksenia@mail.ru, "password": "12345678"}`,
			fasthttp.StatusBadRequest},
		{`xxxxx`, fasthttp.StatusBadRequest},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.input))
		userHandler := CreateDelivery(SetUpDB())
		userHandler.Registration(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}
}

func Test_Register_Success_HttpResponseCode(t *testing.T) {
	tests := []test{
		{`{"name": "Name", "surname": "Surname", "email": "testing1@mail.ru", "password": "password2"}`,
			fasthttp.StatusOK},
	}
	ctx := &fasthttp.RequestCtx{}

	for _, tc := range tests {
		ctx.Request.SetBody(nil)
		ctx.Request.AppendBody([]byte(tc.input))
		userHandler := CreateDelivery(SetUpDB())

		user := new(domain.User)
		if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
			log.Printf("error while unmarshalling JSON: %s", err)
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		//todo убрать deleteByEmail добавить моки
		userHandler.DeleteProfile(ctx)
		userHandler.Registration(ctx)
		assert.Equal(t, tc.want, ctx.Response.Header.StatusCode())
	}
}
