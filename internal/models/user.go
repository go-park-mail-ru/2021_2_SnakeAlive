package models

type Profile struct {
	Id          int    `json:"id"`
	Name        string `json:"name" valid:"required"`
	Surname     string `json:"surname" valid:"required"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type UpdateProfileRequest struct {
	Name        string `json:"name" valid:"required"`
	Surname     string `json:"surname" valid:"required"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Password    string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}
