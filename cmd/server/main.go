package main

import (
	"bytes"
	"context"
	cd "snakealive/m/internal/country/delivery"
	pd "snakealive/m/internal/place/delivery"
	rd "snakealive/m/internal/review/delivery"
	td "snakealive/m/internal/trip/delivery"
	ud "snakealive/m/internal/user/delivery"
	logs "snakealive/m/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/valyala/fasthttp"
)

func SetUpRouter(db *pgxpool.Pool, l *logs.Logger) *router.Router {
	r := router.New()
	r = ud.SetUpUserRouter(db, r, l)
	r = pd.SetUpPlaceRouter(db, r, l)
	r = td.SetUpTripRouter(db, r, l)
	r = rd.SetUpReviewRouter(db, r, l)
	r = cd.SetUpCountryRouter(db, r, l)

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
	l := logs.BuildLogger()

	ctx := ctxzap.ToContext(context.Background(), l.Logger)

	l.Logger.Info("starting server at :8080")

	url := "postgres://tripadvisor:12345@localhost:5432/tripadvisor"
	dbpool, err := pgxpool.Connect(ctx, url)
	if err != nil {
		l.Logger.Fatal("unable to connect to database")
		return
	}
	defer dbpool.Close()

	r := SetUpRouter(dbpool, &l)

	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
		l.Logger.Fatal("failed to start server")
		return
	}
}
