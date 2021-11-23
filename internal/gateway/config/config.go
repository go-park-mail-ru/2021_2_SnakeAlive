package config

import (
	"context"

	logs "snakealive/m/pkg/logger"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	AuthServiceEndpoint  string `envconfig:"AUTH_ENDPOINT" required:"true"`
	TripServiceEndpoint  string `envconfig:"TRIP_ENDPOINT" required:"true"`
	SightServiceEndpoint string `envconfig:"SIGHT_ENDPOINT" required:"true"`
	HTTPPort             string `envconfig:"HTTP_PORT" required:"true"`

	Ctx    context.Context
	Cancel func()
	Logger *zap.Logger
}

func (c *Config) Setup() error {
	if err := envconfig.Process("GATEWAY", c); err != nil {
		return err
	}

	lgr := logs.BuildLogger()
	c.Logger = lgr.Logger
	c.Ctx, c.Cancel = context.WithCancel(context.Background())

	return nil
}
