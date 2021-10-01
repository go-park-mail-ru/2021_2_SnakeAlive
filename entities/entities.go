package entities

import (
	"regexp"
)

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() bool {
	ok, err := regexp.Match(`^\w+[.\w]+@\w+[.\w]+$`, []byte(u.Email))
	if err != nil {
		return false
	}
	if !ok || u.Email == "" {
		return false
	}
	if len(u.Password) < 8 || u.Password == "" || len(u.Password) > 254 || len(u.Email) > 254 {
		return false
	}
	return true
}

type PublicUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Token struct {
	Token string `json:"token,omitempty"`
}

type Place struct {
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
	Photos []string `json:"photos"`
	Author string   `json:"author"`
	Review string   `json:"review"`
}
