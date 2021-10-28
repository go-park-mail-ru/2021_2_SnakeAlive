package userUseCase

import (
	"errors"
	"snakealive/m/domain"
)

func NewUserUseCase(userStorage domain.UserStorage) domain.UserUseCase {
	return userUseCase{userStorage: userStorage}
}

type userUseCase struct {
	userStorage domain.UserStorage
}

func (u userUseCase) Get(key string) (value domain.User, err error) {
	return u.userStorage.Get(key)
}

func (u userUseCase) Add(user domain.User) error {
	return u.userStorage.Add(user)
}

func (u userUseCase) Update(currentUser domain.User, updatedUser domain.User) error {

	user, err := u.Get(updatedUser.Email)
	if err == nil && user.Id != currentUser.Id {
		return errors.New("user with this email already exists") // change later
	}

	return u.userStorage.Update(currentUser.Id, updatedUser)
}

func (u userUseCase) Delete(id int) error {
	return u.userStorage.Delete(id)
}
