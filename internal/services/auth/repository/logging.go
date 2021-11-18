package repository

import (
	"context"

	"go.uber.org/zap"
	"snakealive/m/internal/services/auth/models"
)

const (
	module = "auth_repo"
)

type loggingMiddleware struct {
	logger *zap.Logger

	next AuthRepository
}

func NewLoggingMiddleware(logger *zap.Logger, next AuthRepository) AuthRepository {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (l *loggingMiddleware) GetUserByEmail(ctx context.Context, email string) (u *models.User, err error) {
	l.logger.Info(module,
		zap.Field{
			Key:    "Action",
			String: "GetUserByEmail",
		}, zap.Field{
			Key:       "Request",
			Interface: email,
		},
	)
	defer func() {
		if err != nil {
			l.logger.Info(module,
				zap.Field{
					Key:    "Action",
					String: "GetUserByEmail",
				}, zap.Field{
					Key:       "Request",
					Interface: email,
				},
				zap.Field{
					Key:       "Error",
					Interface: err,
				},
			)
		}
	}()

	return l.next.GetUserByEmail(ctx, email)
}
func (l *loggingMiddleware) GetUserByID(ctx context.Context, ID int64) (u *models.User, err error) {
	l.logger.Info(module,
		zap.Field{
			Key:    "Action",
			String: "GetUserByID",
		}, zap.Field{
			Key:       "Request",
			Interface: ID,
		},
	)
	defer func() {
		if err != nil {
			l.logger.Info(module,
				zap.Field{
					Key:    "Action",
					String: "GetUserByID",
				}, zap.Field{
					Key:       "Request",
					Interface: ID,
				},
				zap.Field{
					Key:       "Error",
					Interface: err,
				},
			)
		}
	}()

	return l.next.GetUserByID(ctx, ID)
}
func (l *loggingMiddleware) CreateUser(ctx context.Context, user *models.User) (u *models.User, err error) {
	l.logger.Info(module,
		zap.Field{
			Key:    "Action",
			String: "CreateUser",
		}, zap.Field{
			Key:       "Request",
			Interface: user,
		},
	)
	defer func() {
		if err != nil {
			l.logger.Info(module,
				zap.Field{
					Key:    "Action",
					String: "CreateUser",
				}, zap.Field{
					Key:       "Request",
					Interface: user,
				},
				zap.Field{
					Key:       "Error",
					Interface: err,
				},
			)
		}
	}()

	return l.next.CreateUser(ctx, user)
}
func (l *loggingMiddleware) UpdateUser(ctx context.Context, user *models.User) (u *models.User, err error) {
	l.logger.Info(module,
		zap.Field{
			Key:    "Action",
			String: "UpdateUser",
		}, zap.Field{
			Key:       "Request",
			Interface: user,
		},
	)
	defer func() {
		if err != nil {
			l.logger.Info(module,
				zap.Field{
					Key:    "Action",
					String: "UpdateUser",
				}, zap.Field{
					Key:       "Request",
					Interface: user,
				},
				zap.Field{
					Key:       "Error",
					Interface: err,
				},
			)
		}
	}()

	return l.next.UpdateUser(ctx, user)
}
func (l *loggingMiddleware) CreateUserSession(ctx context.Context, userID int64, hash string) (err error) {
	l.logger.Info(module,
		zap.Field{
			Key:    "Action",
			String: "CreateUserSession",
		}, zap.Field{
			Key:       "Request",
			Interface: []interface{}{userID, hash},
		},
	)
	defer func() {
		if err != nil {
			l.logger.Info(module,
				zap.Field{
					Key:    "Action",
					String: "CreateUserSession",
				}, zap.Field{
					Key:       "Request",
					Interface: []interface{}{userID, hash},
				},
				zap.Field{
					Key:       "Error",
					Interface: err,
				},
			)
		}
	}()

	return l.next.CreateUserSession(ctx, userID, hash)
}
func (l *loggingMiddleware) ValidateUserSession(ctx context.Context, hash string) (ID int64, err error) {
	l.logger.Info(module,
		zap.Field{
			Key:    "Action",
			String: "ValidateUserSession",
		}, zap.Field{
			Key:       "Request",
			Interface: hash,
		},
	)
	defer func() {
		if err != nil {
			l.logger.Info(module,
				zap.Field{
					Key:    "Action",
					String: "ValidateUserSession",
				}, zap.Field{
					Key:       "Request",
					Interface: hash,
				},
				zap.Field{
					Key:       "Error",
					Interface: err,
				},
			)
		}
	}()

	return l.next.ValidateUserSession(ctx, hash)
}
func (l *loggingMiddleware) RemoveUserSession(ctx context.Context, hash string) (err error) {
	l.logger.Info(module,
		zap.Field{
			Key:    "Action",
			String: "RemoveUserSession",
		}, zap.Field{
			Key:       "Request",
			Interface: hash,
		},
	)
	defer func() {
		if err != nil {
			l.logger.Info(module,
				zap.Field{
					Key:    "Action",
					String: "RemoveUserSession",
				}, zap.Field{
					Key:       "Request",
					Interface: hash,
				},
				zap.Field{
					Key:       "Error",
					Interface: err,
				},
			)
		}
	}()

	return l.next.RemoveUserSession(ctx, hash)
}
