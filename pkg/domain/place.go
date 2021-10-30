package domain

type Place struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Rating      float32  `json:"rating"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}

type Places []Place

type PlaceStorage interface {
	GetById(id int) (value Place, err error)
}

type PlaceUseCase interface {
	GetById(id int) (value Place, err error)
	GetSight(sight Place) (int, []byte)
}
