package http

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func NewLoggingMiddleware(lgr *zap.Logger) Middleware {
	return func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			lgr.Info(string(ctx.Path()),
				zap.String("method", string(ctx.Method())),
				zap.String("remote_addr", string(ctx.RemoteAddr().String())),
				zap.String("url", string(ctx.Path())),
			)

			handler(ctx)

			lgr.Info(string(ctx.Path()),
				zap.Int("status", ctx.Response.StatusCode()),
			)
		}
	}
}
