package userUseCase

import (
	"errors"
	"log"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/domain"

	"github.com/asaskevich/govalidator"
	"github.com/valyala/fasthttp"
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

func (u userUseCase) Validate(user *domain.User) bool {
	if !govalidator.IsEmail(user.Email) ||
		!govalidator.StringLength(user.Password, cnst.MinPasswordLength, cnst.MaxPasswordLength) ||
		!govalidator.MaxStringLength(user.Email, cnst.MaxEmailLength) {
		return false
	}
	return true
}

func (u userUseCase) Login(user *domain.User) (int, error) {
	foundUser, err := u.GetByEmail(user.Email)
	if err != nil {
		log.Printf("error while GetByEmail")
		return fasthttp.StatusNotFound, err
	}

	if foundUser.Password != user.Password {
		return fasthttp.StatusBadRequest, err
	}

	return fasthttp.StatusOK, err
}

func (u userUseCase) Registration(user *domain.User) (int, error) {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		log.Printf("error while validating user")
		return fasthttp.StatusBadRequest, err
	}

	_, err = u.GetByEmail(user.Email)
	if err == nil {
		log.Printf("user with this email already exists")
		return fasthttp.StatusBadRequest, err
	}

	err = u.Add(*user)
	if err == nil {
		log.Printf("error while adding user")
		return fasthttp.StatusBadRequest, err
	}

	return fasthttp.StatusOK, err
}
