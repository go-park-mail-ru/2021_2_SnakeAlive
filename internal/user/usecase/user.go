package userUseCase

import (
	"encoding/json"
	"errors"
	"snakealive/m/internal/domain"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"

	"github.com/asaskevich/govalidator"
	"github.com/microcosm-cc/bluemonday"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func NewUserUseCase(userStorage domain.UserStorage, l *logs.Logger) domain.UserUseCase {
	return userUseCase{
		userStorage: userStorage,
		l:           l,
	}
}

type userUseCase struct {
	userStorage domain.UserStorage
	l           *logs.Logger
}

func (u userUseCase) GetByEmail(key string) (value domain.User, err error) {
	return u.userStorage.GetByEmail(key)
}

func (u userUseCase) GetById(id int) (value domain.User, err error) {
	return u.userStorage.GetById(id)
}

func (u userUseCase) GetPublicById(id int) (value domain.PublicUser, err error) {
	return u.userStorage.GetPublicById(id)
}

func (u userUseCase) Add(user domain.User) error {
	return u.userStorage.Add(user)
}

func (u userUseCase) Update(id int, updatedUser domain.User) error {
	user, err := u.GetByEmail(updatedUser.Email)
	if err != nil && user.Id != id {
		u.l.Logger.Error("user not found")
		return errors.New("failed to get user")
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
		u.l.Logger.Error("error while validating user")
		return false
	}
	return true
}

func (u userUseCase) Login(user *domain.User) (int, error) {
	foundUser, err := u.GetByEmail(user.Email)
	if err != nil {
		u.l.Logger.Error("unable to find user")
		return fasthttp.StatusNotFound, err
	}

	if foundUser.Password != user.Password {
		u.l.Logger.Error("wrong password")
		return fasthttp.StatusBadRequest, err
	}

	return fasthttp.StatusOK, err
}

func (u userUseCase) Registration(user *domain.User) (int, error) {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		u.l.Logger.Error("error while validating user")
		return fasthttp.StatusBadRequest, err
	}

	_, err = u.GetByEmail(user.Email)
	if err == nil {
		u.l.Logger.Error("user with this email already exists")
		return fasthttp.StatusBadRequest, err
	}
	cleanUser := u.SanitizeUser(*user)

	err = u.Add(cleanUser)
	if err != nil {
		u.l.Logger.Error("error while adding user")
		return fasthttp.StatusBadRequest, err
	}

	return fasthttp.StatusOK, err
}

func (u userUseCase) GetProfile(hash string, foundUser domain.User) (int, []byte) {
	response := map[string]string{"name": foundUser.Name, "surname": foundUser.Surname, "email": foundUser.Email,
		"avatar": foundUser.Avatar, "description": foundUser.Description}
	bytes, err := json.Marshal(response)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: %s", zap.Error(err))
		return fasthttp.StatusBadRequest, []byte("{}")
	}
	return fasthttp.StatusOK, bytes
}

func (u userUseCase) UpdateProfile(updatedUser *domain.User, foundUser domain.User, hash string) (int, []byte) {
	_, err := govalidator.ValidateStruct(updatedUser)
	if err != nil {
		u.l.Logger.Error("error while validating user")
		return fasthttp.StatusBadRequest, nil
	}
	cleanUser := u.SanitizeUser(*updatedUser)

	if err = u.Update(foundUser.Id, cleanUser); err != nil {
		u.l.Logger.Error("user with this email already exists")
		return fasthttp.StatusBadRequest, nil
	}

	response := map[string]string{"name": cleanUser.Name, "surname": cleanUser.Surname, "email": cleanUser.Email,
		"avatar": foundUser.Avatar, "description": cleanUser.Description}
	bytes, err := json.Marshal(response)
	if err != nil {
		u.l.Logger.Error("error while marshalling JSON: ", zap.Error(err))
		return fasthttp.StatusBadRequest, nil
	}

	return fasthttp.StatusOK, bytes
}

func (u userUseCase) DeleteProfile(hash string, foundUser domain.User) int {
	u.Delete(foundUser.Id)
	return fasthttp.StatusOK
}

func (u userUseCase) DeleteUserByEmail(user domain.User) int {
	u.userStorage.DeleteByEmail(user)
	return fasthttp.StatusOK
}

func (u userUseCase) AddAvatar(user domain.User, avatar string) (int, []byte) {
	err := u.userStorage.AddAvatar(user.Id, avatar)
	if err != nil {
		return fasthttp.StatusBadRequest, []byte("{}")
	}

	response := map[string]string{"avatar": avatar}
	bytes, err := json.Marshal(response)
	if err != nil {
		return fasthttp.StatusBadRequest, []byte("{}")
	}

	return fasthttp.StatusOK, bytes
}

func (u userUseCase) SanitizeUser(user domain.User) domain.User {
	sanitizer := bluemonday.UGCPolicy()

	user.Name = sanitizer.Sanitize(user.Name)
	user.Surname = sanitizer.Sanitize(user.Surname)
	user.Email = sanitizer.Sanitize(user.Email)
	user.Password = sanitizer.Sanitize(user.Password)
	user.Description = sanitizer.Sanitize(user.Description)
	return user
}
