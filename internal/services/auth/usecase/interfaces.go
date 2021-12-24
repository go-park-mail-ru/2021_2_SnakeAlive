package usecase

import (
	"context"

	"snakealive/m/internal/services/auth/models"
)

type authRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, ID int64) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	CreateUserSession(ctx context.Context, userID int64, hash string) error
	ValidateUserSession(ctx context.Context, hash string) (int64, error)
	RemoveUserSession(ctx context.Context, hash string) error
	GetUserInfo(ctx context.Context, id int) (*models.User, error)
}

type hasher interface {
	EncodeString(value string) string
	DecodeString(value string) (string, error)
}
