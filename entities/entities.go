package entities

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() bool {
	if u.Email == "" {
		return false
	}
	if u.Password == "" {
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
