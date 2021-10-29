package domain

type Place struct {
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
	Photos []string `json:"photos"`
	Author string   `json:"author"`
	Review string   `json:"review"`
}

type Places []Place

type PlaceStorage interface {
	Get(name string) (value Places, exist bool)
}

type PlaceUseCase interface {
	Get(key string) (Places, bool)
}
