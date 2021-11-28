package repository

import (
	"context"

	"snakealive/m/internal/services/auth/models"

	"go.uber.org/zap"
)

const (
	module = "auth_repo"
)

type loggingMiddleware struct {
	logger *zap.SugaredLogger

	next AuthRepository
}

func NewLoggingMiddleware(logger *zap.SugaredLogger, next AuthRepository) AuthRepository {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (l *loggingMiddleware) GetUserByEmail(ctx context.Context, email string) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "GetUserByEmail",
		"Request", email,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUserByEmail",
				"Request", email,
				"Error", err,
			)
		}
	}()

	return l.next.GetUserByEmail(ctx, email)
}
func (l *loggingMiddleware) GetUserByID(ctx context.Context, ID int64) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "GetUserByID",
		"Request", ID,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUserByID",
				"Request", ID,
				"Error", err,
			)
		}
	}()

	return l.next.GetUserByID(ctx, ID)
}
func (l *loggingMiddleware) CreateUser(ctx context.Context, user *models.User) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "CreateUser",
		"Request", user,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "CreateUser",
				"Request", user,
				"Error", err,
			)
		}
	}()

	return l.next.CreateUser(ctx, user)
}
func (l *loggingMiddleware) UpdateUser(ctx context.Context, user *models.User) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "UpdateUser",
		"Request", user,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateUser",
				"Request", user,
				"Error", err,
			)
		}
	}()

	return l.next.UpdateUser(ctx, user)
}
func (l *loggingMiddleware) CreateUserSession(ctx context.Context, userID int64, hash string) (err error) {
	l.logger.Infow(module,
		"Action", "CreateUserSession",
		"Request", userID,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "CreateUserSession",
				"Request", userID,
				"Error", err,
			)
		}
	}()

	return l.next.CreateUserSession(ctx, userID, hash)
}
func (l *loggingMiddleware) ValidateUserSession(ctx context.Context, hash string) (ID int64, err error) {
	l.logger.Infow(module,
		"Action", "ValidateUserSession",
		"Request", hash,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "ValidateUserSession",
				"Request", hash,
				"Error", err,
			)
		}
	}()

	return l.next.ValidateUserSession(ctx, hash)
}
func (l *loggingMiddleware) RemoveUserSession(ctx context.Context, hash string) (err error) {
	l.logger.Infow(module,
		"Action", " RemoveUserSession",
		"Request", hash,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", " RemoveUserSession",
				"Request", hash,
				"Error", err,
			)
		}
	}()

	return l.next.RemoveUserSession(ctx, hash)
}

func (l *loggingMiddleware) GetUserInfo(ctx context.Context, id int) (u *models.User, err error) {
	l.logger.Infow(module,
		"Action", "GetUserByEmail",
		"Request", id,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUserInfo",
				"Request", id,
				"Error", err,
			)
		}
	}()

	return l.next.GetUserInfo(ctx, id)
}
