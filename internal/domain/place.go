package domain

import "github.com/valyala/fasthttp"

type Place struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Lat         float32  `json:"lat"`
	Lng         float32  `json:"lng"`
	Rating      float32  `json:"rating"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}

type TopPlace struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Tags    []string `json:"tags"`
	Photos  []string `json:"photos"`
	Country string   `json:"country"`
	Rating  float32  `json:"rating"`
}

type PlaceCoord struct {
	Id  int     `json:"id"`
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type PlaceCoords []PlaceCoord

type TopPlaces []TopPlace

type PlaceHandler interface {
	PlacesByCountry(ctx *fasthttp.RequestCtx)
	Place(ctx *fasthttp.RequestCtx)
}

type PlaceStorage interface {
	GetById(id int) (value Place, err error)
	GetPlacesByCountry(value string) (TopPlaces, error)
}

type PlaceUseCase interface {
	GetById(id int) (value Place, err error)
	GetSight(sight Place) (int, []byte)
	GetPlacesByCountry(value string) ([]byte, error)
}
