package domain

type User struct {
	Name     string `json:"name" valid:"required,alpha"`
	Surname  string `json:"surname" valid:"required,alpha"`
	Email    string `json:"email" valid:"required,email,maxstringlength(254)"`
	Password string `json:"password" valid:"required,stringlength(8|254)"`
}

type UserStorage interface {
	Add(key string, value User)
	Get(key string) (value User, exist bool)
	Delete(key string)
	Update(key string, value User)
}

type UserUseCase interface {
	Get(key string) (User, bool)
	Add(user User)
	Delete(key string)
	Update(currentUser User, updatedUser User) bool
}
