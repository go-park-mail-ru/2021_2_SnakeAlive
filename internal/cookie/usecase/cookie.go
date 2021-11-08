package cookieUseCase

import "snakealive/m/internal/domain"

func NewCookieUseCase(cookieStorage domain.CookieStorage) domain.CookieUseCase {
	return cookieUseCase{cookieStorage: cookieStorage}
}

type cookieUseCase struct {
	cookieStorage domain.CookieStorage
}

func (c cookieUseCase) Add(key string, userId int) error {
	return c.cookieStorage.Add(key, userId)
}

func (c cookieUseCase) Get(value string) (user domain.User, err error) {
	return c.cookieStorage.Get(value)
}

func (c cookieUseCase) Delete(value string) error {
	return c.cookieStorage.Delete(value)
}
