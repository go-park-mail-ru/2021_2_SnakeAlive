package models

type User struct {
	ID          int64
	Name        string
	Surname     string
	Email       string
	Password    string
	Image       string
	Description string
}

type Session struct {
	Cookie string
	Token  string
}
