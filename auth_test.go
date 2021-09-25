package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

// serve serves http request using provided fasthttp handler
func serve(handler fasthttp.RequestHandler, req *http.Request) (*http.Response, error) {
	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	go func() {
		err := fasthttp.Serve(ln, handler)
		if err != nil {
			panic(fmt.Errorf("failed to serve: %v", err))
		}
	}()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return ln.Dial()
			},
		},
	}

	return client.Do(req)
}

type test struct {
	input  []byte
	method string
	want   int
}

func TestLogin(t *testing.T) {
	tests := []test{
		{[]byte(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "password": "password"}`), "POST",
			fasthttp.StatusOK},
		{[]byte(`{"name": "name2", "surname": "surname2", "email": "xxxxxx@mail.ru", "password": "pass"}`), "POST",
			fasthttp.StatusNotFound},
		{[]byte(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "password": "xxxxxxxx"}`), "POST",
			fasthttp.StatusBadRequest},
		{[]byte(`xxxxx`), "POST", fasthttp.StatusBadRequest},
	}

	for _, tc := range tests {
		t.Log(tc.input, tc.want)
		r, err := http.NewRequest(tc.method, "http://:8080", bytes.NewBuffer(tc.input))
		if err != nil {
			t.Error(err)
		}
		res, err := serve(login, r)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tc.want, res.StatusCode)
	}
}

func TestRegister(t *testing.T) {
	tests := []test{
		{([]byte(`{"name": "name2", "surname": "surname2", "email": "asdf@mail.ru", "password": "pass"}`)), "POST",
			fasthttp.StatusOK},
		{[]byte(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "password": "pass"}`), "POST",
			fasthttp.StatusBadRequest},
		{[]byte(`{"name": "name2", "surname": "surname2", "email": "alex@mail.ru", "xxxxx": "xxxxxxxx"}`), "POST",
			fasthttp.StatusBadRequest},
		{[]byte(`{"name": "name2", "surname": "surname2", "xxxx": "alex@mail.ru", "password": "xxxxxxxx"}`), "POST",
			fasthttp.StatusBadRequest},
		{[]byte(`xxxxx`), "POST", fasthttp.StatusBadRequest},
	}

	for _, tc := range tests {
		t.Log(tc.input, tc.want)
		r, err := http.NewRequest(tc.method, "http://:8080", bytes.NewBuffer(tc.input))
		if err != nil {
			t.Error(err)
		}
		res, err := serve(registration, r)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tc.want, res.StatusCode)
	}
}
