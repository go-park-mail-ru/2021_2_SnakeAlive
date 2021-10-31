package domain

type Place struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Rating      int      `json:"rating"`
	Description string   `json:"sescription"`
	Tags        []string `json:"tags"`
}

type Places []Place

type PlaceStorage interface {
	Get(name string) (value Place, exist bool)
}

type PlaceUseCase interface {
	GetById(id int) (value Place, err error)
	GetSight(sight Place) (int, []byte)
	GetPlacesByCountry(value string) ([]byte, error)
}
