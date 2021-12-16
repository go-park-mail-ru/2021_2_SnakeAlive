package models

type Place struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Rating      float32  `json:"rating"`
	Tags        []Tag    `json:"tags"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Day         int      `json:"day"`
}

type Trip struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Sights      []Place `json:"sights"`
	Albums      []Album `json:"albums"`
	Users       []int   `json:"users"`
}

type TripSight struct {
	Id  int     `json:"id"`
	Lng float32 `json:"lng"`
	Lat float32 `json:"lat"`
}

type Album struct {
	Id          int      `json:"id"`
	TripId      int      `json:"trip_id"`
	UserId      int      `json:"user_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}

type TripUser struct {
	Email string `json:"email"`
}

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name" valid:"required"`
	Surname string `json:"surname" valid:"required"`
	Avatar  string `json:"avatar"`
}
type TripWithUserInfo struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Sights      []Place    `json:"sights"`
	Albums      []Album    `json:"albums"`
	Users       []UserInfo `json:"users"`
}
