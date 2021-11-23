package repository

import (
	"context"

	"go.uber.org/zap"
	"snakealive/m/internal/services/sight/models"
)

const (
	module = "sight_repo"
)

type loggingMiddleware struct {
	logger *zap.SugaredLogger

	next SightRepository
}

func (l *loggingMiddleware) GetSightsByCountry(ctx context.Context, country string) (sights []models.Sight, err error) {
	l.logger.Infow(module,
		"Action", "GetSightsByCountry",
		"Request", country,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetSightsByCountry",
				"Request", country,
				"Error", err,
			)
		}
	}()

	return l.next.GetSightsByCountry(ctx, country)
}

func (l *loggingMiddleware) GetSightByID(ctx context.Context, id int) (sight models.Sight, err error) {
	l.logger.Infow(module,
		"Action", "GetSightByID",
		"Request", id,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetSightByID",
				"Request", id,
				"Error", err,
			)
		}
	}()

	return l.next.GetSightByID(ctx, id)
}

func NewLoggingMiddleware(
	logger *zap.SugaredLogger,

	next SightRepository,
) SightRepository {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}
