package main

import (
	"bytes"
	"fmt"
	"snakealive/m/delivery"
	"snakealive/m/domain"
	"snakealive/m/usecase"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Router(ss delivery.SessionServer) *router.Router {

	r := router.New()
	r.POST("/login", ss.Login)
	r.POST("/register", ss.Registration)
	// r.GET("/country/{name}", delivery.PlacesList)
	// r.GET("/profile", delivery.Profile)
	// r.DELETE("/logout", delivery.Logout)
	return r
}

func corsMiddleware(handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://194.58.104.204") // set domain
		ctx.Response.Header.Set("Content-Type", "application/json; charset=utf8")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type")
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

	//userhttp.NewUserHandler(e, usecase.NewUserUsecase(repository.NewUserMemoryRepository()))

	r := Router(delivery.NewSessionServer(usecase.NewUseCase(domain.NewUserStorage())))

	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
		fmt.Println("failed to start server:", err)
		return
	}
}
