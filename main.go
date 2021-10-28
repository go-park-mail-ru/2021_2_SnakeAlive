package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	pd "snakealive/m/place/delivery"
	pr "snakealive/m/place/repository"
	pu "snakealive/m/place/usecase"
	ud "snakealive/m/user/delivery"
	ur "snakealive/m/user/repository"
	uu "snakealive/m/user/usecase"

	"github.com/fasthttp/router"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/valyala/fasthttp"
)

func Router(db *pgxpool.Pool) *router.Router {
	r := router.New()

	ud.NewUserHandler(uu.NewUserUseCase(ur.NewUserStorage(db)), r)
	pd.NewPlaceHandler(pu.NewPlaceUseCase(pr.NewPlaceStorage()), r)

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

	url := "postgres://tripadvisor:12345@localhost:5432/tripadvisor"
	dbpool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	r := Router(dbpool)

	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
		fmt.Println("failed to start server:", err)
		return
	}
}
