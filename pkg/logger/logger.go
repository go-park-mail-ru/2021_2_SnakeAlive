package logs

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Logger *zap.Logger
}

func BuildLogger() Logger {
	var zapCfg = zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err := zapCfg.Build()
	if err != nil {
		panic(err)
	}

	var logs Logger
	logs.Logger = log
	return logs
}

func AccessLogMiddleware(l *Logger, handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		l.Logger.Info(string(ctx.Path()),
			zap.String("method", string(ctx.Method())),
			zap.String("remote_addr", string(ctx.RemoteAddr().String())),
			zap.String("url", string(ctx.Path())),
		)

		handler(ctx)

		l.Logger.Info(string(ctx.Path()),
			zap.Int("status", ctx.Response.StatusCode()),
		)
	}
}
