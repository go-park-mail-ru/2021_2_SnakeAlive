package domain

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" valid:"required,alpha"`
	Surname  string `json:"surname" valid:"required,alpha"`
	Email    string `json:"email" valid:"required,email,maxstringlength(254)"`
	Password string `json:"password" valid:"required,stringlength(8|254)"`
}

type UserStorage interface {
	Add(value User) error
	Get(key string) (value User, err error)
	Delete(id int) error
	Update(id int, value User) error
}

type UserUseCase interface {
	Add(user User) error
	Get(key string) (value User, err error)
	Delete(id int) error
	Update(currentUser User, updatedUser User) error
}
