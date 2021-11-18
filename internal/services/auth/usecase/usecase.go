package usecase

import (
	"context"

	"github.com/twinj/uuid"
	"snakealive/m/internal/services/auth/models"
	"snakealive/m/pkg/errors"
)

type AuthUseCase interface {
	LoginUser(ctx context.Context, user *models.User) (models.Session, error)
	LogoutUser(ctx context.Context, session string) error
	ValidateSession(ctx context.Context, session string) (int64, error)

	RegisterUser(ctx context.Context, user *models.User) (models.Session, error)
	GetUser(ctx context.Context, ID int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
}

type authUseCase struct {
	hashGenerator hasher
	repo          authRepository
	uuidGen       uuid.UUID
}

func (a *authUseCase) LoginUser(ctx context.Context, user *models.User) (models.Session, error) {
	repoUser, err := a.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return models.Session{}, err
	}

	pass, _ := a.hashGenerator.DecodeString(repoUser.Password)
	if pass != user.Password {
		return models.Session{}, errors.WrongUserPassword
	}

	cookie := a.uuidGen.String()
	if err = a.repo.CreateUserSession(ctx, repoUser.ID, cookie); err != nil {
		return models.Session{}, err
	}

	return models.Session{
		Cookie: cookie,
		Token:  "??",
	}, nil
}

func (a *authUseCase) LogoutUser(ctx context.Context, session string) error {
	return a.repo.RemoveUserSession(ctx, session)
}

func (a *authUseCase) ValidateSession(ctx context.Context, session string) (int64, error) {
	return a.repo.ValidateUserSession(ctx, session)
}

func (a *authUseCase) RegisterUser(ctx context.Context, user *models.User) (models.Session, error) {
	user.Password = a.hashGenerator.EncodeString(user.Password)
	user, err := a.repo.CreateUser(ctx, user)
	if err != nil {
		return models.Session{}, err
	}

	cookie := a.uuidGen.String()
	if err = a.repo.CreateUserSession(ctx, user.ID, cookie); err != nil {
		return models.Session{}, err
	}

	return models.Session{
		Cookie: cookie,
		Token:  "??",
	}, nil
}

func (a *authUseCase) GetUser(ctx context.Context, ID int64) (*models.User, error) {
	return a.repo.GetUserByID(ctx, ID)
}

func (a *authUseCase) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if user.Password != "" {
		user.Password = a.hashGenerator.EncodeString(user.Password)
	}

	return a.repo.UpdateUser(ctx, user)
}

func NewAuthUseCase(
	hashGenerator hasher,
	repo authRepository,
) AuthUseCase {
	return &authUseCase{
		hashGenerator: hashGenerator,
		repo:          repo,
		uuidGen:       uuid.NewV4(),
	}
}
