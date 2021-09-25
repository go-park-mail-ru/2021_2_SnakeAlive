package entities

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Place struct {
	Name   string
	Tags   []string
	Photos []string
	Author string
	Review string
}
