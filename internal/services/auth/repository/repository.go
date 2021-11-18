package repository

import (
	"context"
	"database/sql"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	"snakealive/m/internal/services/auth/models"
	"snakealive/m/pkg/errors"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, ID int64) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	CreateUserSession(ctx context.Context, userID int64, hash string) error
	ValidateUserSession(ctx context.Context, hash string) (int64, error)
	RemoveUserSession(ctx context.Context, hash string) error
}

type authRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (a *authRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := a.queryFactory.CreateUpdateUser(user)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	if err := row.Scan(
		&user.Name, &user.Surname, &user.Email,
		&user.Image, &user.Description,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := a.queryFactory.CreateGetUserByEmail(email)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return user, nil
}

func (a *authRepository) GetUserByID(ctx context.Context, ID int64) (*models.User, error) {
	query := a.queryFactory.CreateGetUserByID(ID)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	user := &models.User{}
	if err := row.Scan(
		&user.Name, &user.Surname, &user.Email,
		&user.Image, &user.Description,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return user, nil
}

func (a *authRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := a.queryFactory.CreateCreateUser(user)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	if err := row.Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authRepository) CreateUserSession(ctx context.Context, userID int64, hash string) error {
	query := a.queryFactory.CreateCreateUserSession(userID, hash)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		return err
	}

	return nil
}

func (a *authRepository) ValidateUserSession(ctx context.Context, hash string) (int64, error) {
	userID := int64(0)
	query := a.queryFactory.CreateValidateUserSession(hash)
	if err := a.conn.QueryRow(ctx, query.Request, query.Params...).Scan(&userID); err != nil {
		if err == sql.ErrNoRows {
			return userID, errors.SessionDoesNotExist
		}

		return userID, err
	}

	return userID, nil
}

func (a *authRepository) RemoveUserSession(ctx context.Context, hash string) error {
	query := a.queryFactory.CreateRemoveUserSession(hash)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		return err
	}

	return nil
}

func NewAuthRepostory() AuthRepository {
	return &authRepository{}
}
