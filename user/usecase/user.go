package userUseCase

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

func (u userUseCase) Update(currentUser domain.User, updatedUser domain.User) bool {
	/*
		user, err := u.Get(updatedUser.Email)
		if !err && user.id != currentUser.id {
			return false
		}
	*/

	u.userStorage.Update(currentUser.Email, updatedUser)
	return true
}

func (u userUseCase) Delete(key string) {
	u.userStorage.Delete(key)
}
