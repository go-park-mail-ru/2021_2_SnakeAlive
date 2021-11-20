package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	logs "snakealive/m/pkg/logger"
)

type Config struct {
	AuthServiceEndpoint string `envconfig:"AUTH_ENDPOINT" required:"true"`
	HTTPPort            string `envconfig:"HTTP_PORT" required:"true"`

	Ctx    context.Context
	Cancel func()
	Logger *zap.Logger
}

func (c Config) Setup() error {
	if err := envconfig.Process("GATEWAY", &c); err != nil {
		return err
	}

	lgr := logs.BuildLogger()
	c.Logger = lgr.Logger
	c.Ctx, c.Cancel = context.WithCancel(context.Background())

	return nil
}
