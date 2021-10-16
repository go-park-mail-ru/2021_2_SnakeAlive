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
	fmt.Println("GET key = ", key, "res = ", val)
	return u.userStorage.Get(key)
}

func (u userUsecase) Add(key string, value domain.User) {
	fmt.Println("Add key = ", key, "res = ", value)
	u.userStorage.Add(key, value)
}
