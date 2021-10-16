package usecase

import (
	"snakealive/m/domain"
)

func NewUserUseCase(userStorage domain.UserStorage) domain.UserUseCase {
	return userUseCase{userStorage: userStorage}
}

type userUseCase struct {
	userStorage domain.UserStorage
}

func (u userUseCase) Get(key string) (domain.User, bool) {
	return u.userStorage.Get(key)
}

func (u userUseCase) Add(user domain.User) {
	u.userStorage.Add(user.Email, user)
}
