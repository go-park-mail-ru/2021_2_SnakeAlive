package http

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/valyala/fasthttp"
	"snakealive/m/entities"
	mock_storage "snakealive/m/session/storage/mock"
)

func Test(t *testing.T) {
	ctrl := gomock.NewController(t)
	storage := mock_storage.NewMockUserStorage(ctrl)
	server := NewSessionServer(storage)

	u := entities.User{
		Name:     "123",
		Surname:  "123",
		Email:    "123",
		Password: "123",
	}
	bytes, _ := json.Marshal(u)

	ctx := &fasthttp.RequestCtx{}
	ctx.Request.SetBody(bytes)
	storage.EXPECT().Get("123").Return(u, true)
	server.Login(ctx)

	fmt.Println(ctx.Response.StatusCode())
}
