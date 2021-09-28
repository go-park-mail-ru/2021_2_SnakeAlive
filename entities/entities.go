package entities

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserJSON struct {
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
