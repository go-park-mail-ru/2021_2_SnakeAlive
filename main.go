package main

import (
	"bytes"
	"fmt"
	pd "snakealive/m/place/delivery"
	pr "snakealive/m/place/repository"
	pu "snakealive/m/place/usecase"
	ud "snakealive/m/user/delivery"
	ur "snakealive/m/user/repository"
	uu "snakealive/m/user/usecase"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Router(user interface{}, place interface{}) *router.Router {
	r := router.New()
	userHandler := user.(ud.UserHandler)
	placeHandler := place.(pd.PlaceHandler)
	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Registration)
	r.GET("/country/{name}", placeHandler.PlacesList)
	r.GET("/profile", userHandler.Profile)
	r.DELETE("/logout", userHandler.Logout)
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
	userLayer := ud.NewUserHandler(uu.NewUserUseCase(ur.NewUserStorage()))
	placeLayer := pd.NewPlaceHandler(pu.NewPlaceUseCase(pr.NewPlaceStorage()))
	r := Router(userLayer, placeLayer)

	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
		fmt.Println("failed to start server:", err)
		return
	}
}
