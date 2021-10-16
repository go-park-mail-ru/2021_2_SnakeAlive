package domain

type Place struct {
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
	Photos []string `json:"photos"`
	Author string   `json:"author"`
	Review string   `json:"review"`
}
