package logs

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func BuildLogger() {
	var zapCfg = zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err := zapCfg.Build()
	if err != nil {
		panic(err)
	}

	logger = log
}

func GetLogger() *zap.Logger {
	return logger
}

func AccessLogMiddleware(handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		logger.Info(string(ctx.Path()),
			zap.String("method", string(ctx.Method())),
			zap.String("remote_addr", string(ctx.RemoteAddr().String())),
			zap.String("url", string(ctx.Path())),
		)

		handler(ctx)

		logger.Info(string(ctx.Path()),
			zap.Int("status", ctx.Response.StatusCode()),
		)
	}
}
