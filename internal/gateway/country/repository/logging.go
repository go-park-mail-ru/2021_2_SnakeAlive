package repository

import (
	"context"

	"go.uber.org/zap"
	"snakealive/m/internal/models"
)

const (
	module = "country_repository"
)

type loggingMiddleware struct {
	logger *zap.SugaredLogger

	next CountryStorage
}

func (l *loggingMiddleware) GetCountriesList(ctx context.Context) (c models.Countries, err error) {
	l.logger.Infow(module,
		"Action", "GetCountriesList",
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetCountriesList",
				"Error", err,
			)
		}
	}()

	return l.next.GetCountriesList(ctx)
}

func (l *loggingMiddleware) GetById(ctx context.Context, id int) (c models.Country, err error) {
	l.logger.Infow(module,
		"Action", "GetById",
		"Request", id,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetById",
				"Request", id,
				"Error", err,
			)
		}
	}()

	return l.next.GetById(ctx, id)
}

func (l *loggingMiddleware) GetByName(ctx context.Context, name string) (c models.Country, err error) {
	l.logger.Infow(module,
		"Action", "GetByName",
		"Request", name,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetByName",
				"Request", name,
				"Error", err,
			)
		}
	}()

	return l.next.GetByName(ctx, name)
}

func NewLoggingMiddleware(lgr *zap.SugaredLogger, next CountryStorage) CountryStorage {
	return &loggingMiddleware{
		logger: lgr,
		next:   next,
	}
}
