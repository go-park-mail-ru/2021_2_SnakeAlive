package usecase

import (
	"fmt"
	"snakealive/m/domain"
)

func NewUseCase(userStorage domain.UserStorage) domain.Usecase {
	return userUsecase{userStorage: userStorage}
}

type userUsecase struct {
	userStorage domain.UserStorage
}

func (u userUsecase) Get(key string) (domain.User, bool) {
	val, _ := u.userStorage.Get(key)
	fmt.Println("key = ", key, "res = ", val)
	return u.userStorage.Get(key)
}
