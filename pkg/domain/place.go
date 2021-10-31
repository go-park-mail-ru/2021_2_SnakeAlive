package domain

type Place struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Rating      float32  `json:"rating"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}

type TopPlace struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
	Author string   `json:"author"`
	Review string   `json:"review"`
}

type TopPlaces []TopPlace

type PlaceStorage interface {
	GetById(id int) (value Place, err error)
	GetPlacesByCountry(value string) (TopPlaces, error)
}

type PlaceUseCase interface {
	GetById(id int) (value Place, err error)
	GetSight(sight Place) (int, []byte)
	GetPlacesByCountry(value string) ([]byte, error)
}
