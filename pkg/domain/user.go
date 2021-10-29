package domain

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" valid:"required"`
	Surname  string `json:"surname" valid:"required"`
	Email    string `json:"email" valid:"required,email,maxstringlength(254)"`
	Password string `json:"password" valid:"required,stringlength(8|254)"`
}

type UserStorage interface {
	Add(value User) error
	GetById(id int) (value User, err error)
	GetByEmail(key string) (value User, err error)
	Delete(id int) error
	Update(id int, value User) error
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
}
