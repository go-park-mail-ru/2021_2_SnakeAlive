package cookieUseCase

import (
	"snakealive/m/internal/domain"
	logs "snakealive/m/pkg/logger"
)

func NewCookieUseCase(cookieStorage domain.CookieStorage, l *logs.Logger) domain.CookieUseCase {
	return cookieUseCase{
		cookieStorage: cookieStorage,
		l:             l,
	}
}

type cookieUseCase struct {
	l             *logs.Logger
	cookieStorage domain.CookieStorage
}

func (c cookieUseCase) Add(key string, userId int) error {
	return c.cookieStorage.Add(key, userId)
}

func (c cookieUseCase) Get(value string) (user domain.User, err error) {
	user, err = c.cookieStorage.Get(value)
	if err != nil {
		c.l.Logger.Error("unable to find cookie")
	}
	return user, err
}

func (c cookieUseCase) Delete(value string) error {
	return c.cookieStorage.Delete(value)
}
