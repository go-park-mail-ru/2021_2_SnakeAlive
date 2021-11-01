package logs

import (
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
