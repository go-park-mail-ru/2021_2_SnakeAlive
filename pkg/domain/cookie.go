package domain

type CookieStorage interface {
	Add(key string, userId int) error
	Get(value string) (user User, err error)
	Delete(value string) error
}

type CookieUseCase interface {
	Add(key string, userId int) error
	Get(value string) (user User, err error)
	Delete(value string) error
}
