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

func (l *loggingMiddleware) GetTags(ctx context.Context) (tags []models.Tag, err error) {
	l.logger.Infow(module,
		"Action", "GetTags",
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetTags",
				"Error", err,
			)
		}
	}()

	return l.next.GetTags(ctx)
}

func (l *loggingMiddleware) GetSightByTag(ctx context.Context, tag int64) (sights []models.Sight, err error) {
	l.logger.Infow(module,
		"Action", "GetSightByTag",
		"Request", tag,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetSightByTag",
				"Request", tag,
				"Error", err,
			)
		}
	}()

	return l.next.GetSightByTag(ctx, tag)
}

func (l *loggingMiddleware) GetSightByIDs(ctx context.Context, ids []int64) (sights []models.Sight, err error) {
	l.logger.Infow(module,
		"Action", "GetSightByIDs",
		"Request", ids,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetSightByIDs",
				"Request", ids,
				"Error", err,
			)
		}
	}()

	return l.next.GetSightByIDs(ctx, ids)
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

func (l *loggingMiddleware) SearchSights(ctx context.Context, search string, skip, limit int64) (sights []models.Sight, err error) {
	l.logger.Infow(module,
		"Action", "GetSightByID",
		"Request", []interface{}{search, skip, limit},
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetSightByID",
				"Request", []interface{}{search, skip, limit},
				"Error", err,
			)
		}
	}()

	return l.next.SearchSights(ctx, search, skip, limit)
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
