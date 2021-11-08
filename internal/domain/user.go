package domain

import "github.com/valyala/fasthttp"

//go:generate mockgen -source=user.go -destination=/mocks/mock.go
//mockgen -source=user.go -destination=mocks/mock_doer.go -package=mocks

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" valid:"required"`
	Surname  string `json:"surname" valid:"required"`
	Email    string `json:"email" valid:"required,email,maxstringlength(254)"`
	Password string `json:"password" valid:"required,stringlength(8|254)"`
	Avatar   string `json:"avatar"`
}

type PublicUser struct {
	Id      int    `json:"id"`
	Name    string `json:"name" valid:"required"`
	Surname string `json:"surname" valid:"required"`
	Avatar  string `json:"avatar"`
}

type UserHandler interface {
	Login(ctx *fasthttp.RequestCtx)
	Registration(ctx *fasthttp.RequestCtx)
	Logout(ctx *fasthttp.RequestCtx)
	GetProfile(ctx *fasthttp.RequestCtx)
	UpdateProfile(ctx *fasthttp.RequestCtx)
	DeleteProfile(ctx *fasthttp.RequestCtx)
	DeleteProfileByEmail(ctx *fasthttp.RequestCtx)
	UploadAvatar(ctx *fasthttp.RequestCtx)
}

type UserStorage interface {
	Add(value User) error
	GetById(id int) (value User, err error)
	GetPublicById(id int) (value PublicUser, err error)
	GetByEmail(key string) (value User, err error)
	Delete(id int) error
	Update(id int, value User) error
	DeleteByEmail(user User) error
	AddAvatar(id int, avatar string) error
}

type UserUseCase interface {
	Add(user User) error
	GetById(id int) (value User, err error)
	GetPublicById(id int) (value PublicUser, err error)
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
	SanitizeUser(user User) User
}
