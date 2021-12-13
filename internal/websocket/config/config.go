package config

import (
	"context"

	logs "snakealive/m/pkg/logger"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	DBUrl string `envconfig:"DB_URL" required:"true"`
	Port  string `envconfig:"PORT" required:"true"`

	Ctx    context.Context
	Cancel func()
	Logger *zap.Logger
}

func (c *Config) Setup() error {
	if err := envconfig.Process("WEBSOCKET", c); err != nil {
		return err
	}

	lgr := logs.BuildLogger()
	c.Logger = lgr.Logger
	c.Ctx, c.Cancel = context.WithCancel(context.Background())

	return nil
}
