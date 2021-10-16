package domain

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStorage interface {
	Add(key string, value User)
	Get(key string) (value User, exist bool)
	Delete(key string)
}

type UserUseCase interface {
	Get(key string) (User, bool)
	Add(user User)
}
