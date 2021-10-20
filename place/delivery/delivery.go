package placeDelivery

import (
	"encoding/json"
	"log"
	"snakealive/m/domain"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const CookieName = "SnakeAlive"

type PlaceHandler interface {
	PlacesList(ctx *fasthttp.RequestCtx)
}

type placeHandler struct {
	PlaceUseCase domain.PlaceUseCase
}

func NewPlaceHandler(PlaceUseCase domain.PlaceUseCase, r *router.Router) {
	placeHandler := placeHandler{PlaceUseCase: PlaceUseCase}

	r.GET("/country/{name}", placeHandler.PlacesList)
}

func (s *placeHandler) PlacesList(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("name").(string)
	if _, found := s.PlaceUseCase.Get(param); !found {
		log.Printf("country not found")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	response, _ := s.PlaceUseCase.Get(param)
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}
