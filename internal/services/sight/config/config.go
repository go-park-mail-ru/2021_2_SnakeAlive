package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	logs "snakealive/m/pkg/logger"
)

type Config struct {
	DBUrl    string `envconfig:"DB_URL" required:"true"`
	GRPCPort string `envconfig:"GRPC_PORT" required:"true"`

	Ctx    context.Context
	Cancel func()
	Logger *zap.Logger
}

func (c *Config) Setup() error {
	if err := envconfig.Process("SIGHT", c); err != nil {
		return err
	}

	lgr := logs.BuildLogger()
	c.Logger = lgr.Logger
	c.Ctx, c.Cancel = context.WithCancel(context.Background())

	return nil
}
