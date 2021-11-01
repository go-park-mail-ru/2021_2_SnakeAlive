package domain

//go:generate mockgen -source=user.go -destination=/mocks/mock.go

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" valid:"required"`
	Surname  string `json:"surname" valid:"required"`
	Email    string `json:"email" valid:"required,email,maxstringlength(254)"`
	Password string `json:"password" valid:"required,stringlength(8|254)"`
	Avatar   string `json:"avatar"`
}

type UserStorage interface {
	Add(value User) error
	GetById(id int) (value User, err error)
	GetByEmail(key string) (value User, err error)
	Delete(id int) error
	Update(id int, value User) error
	DeleteByEmail(user User) error
	AddAvatar(id int, avatar string) error
}

type UserUseCase interface {
	Add(user User) error
	GetById(id int) (value User, err error)
	GetByEmail(key string) (value User, err error)
	Delete(id int) error
	Update(id int, updatedUser User) error
	Validate(user *User) bool
	Login(user *User) (int, error)
	Registration(user *User) (int, error)
	GetProfile(hash string, user User) (int, []byte)
	UpdateProfile(updatedUser *User, foundUser User, hash string) (int, []byte)
	DeleteProfile(hash string, foundUser User) int
	DeleteUserByEmail(user User) int
	AddAvatar(user User, avatar string) error
}
