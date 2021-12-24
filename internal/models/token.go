package models

type Session struct {
	Token  string `json:"token"`
	Cookie string `json:"-"`
}
