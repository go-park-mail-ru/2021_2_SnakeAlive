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

func (u userUseCase) GetByEmail(key string) (value domain.User, err error) {
	return u.userStorage.GetByEmail(key)
}

func (u userUseCase) GetById(id int) (value domain.User, err error) {
	return u.userStorage.GetById(id)
}

func (u userUseCase) Add(user domain.User) error {
	return u.userStorage.Add(user)
}

func (u userUseCase) Update(id int, updatedUser domain.User) error {

	user, err := u.GetByEmail(updatedUser.Email)
	if err == nil && user.Id != id {
		return errors.New("user with this email already exists") // change later
	}

	return u.userStorage.Update(id, updatedUser)
}

func (u userUseCase) Delete(id int) error {
	return u.userStorage.Delete(id)
}
