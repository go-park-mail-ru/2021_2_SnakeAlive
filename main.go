package main

import (
	"bytes"
	"fmt"
	"snakealive/m/http"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Router() *router.Router {
	r := router.New()

	r.POST("/login", http.Login)
	r.POST("/register", http.Registration)
	r.GET("/country/{name}", http.PlacesList)
	r.GET("/profile", http.Profile)
	r.GET("/logout", http.Logout)
	return r
}

func corsMiddleware(handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*") // set domain
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

func main() {
	fmt.Println("starting server at :8080")

	r := Router()

	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
		fmt.Println("failed to start server:", err)
		return
	}
}
