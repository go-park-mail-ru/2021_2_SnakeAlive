package delivery

import (
	"encoding/json"
	"fmt"
	"snakealive/m/internal/gateway/trip/usecase"
	"snakealive/m/internal/services/trip/models"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
	"strconv"

	"github.com/valyala/fasthttp"
)

type TripGatewayDelivery interface {
	Trip(ctx *fasthttp.RequestCtx)
	AddTrip(ctx *fasthttp.RequestCtx)
	UpdateTrip(ctx *fasthttp.RequestCtx)
	DeleteTrip(ctx *fasthttp.RequestCtx)
	Album(ctx *fasthttp.RequestCtx)
	AddAlbum(ctx *fasthttp.RequestCtx)
	UpdateAlbum(ctx *fasthttp.RequestCtx)
	DeleteAlbum(ctx *fasthttp.RequestCtx)
	UploadPhoto(ctx *fasthttp.RequestCtx)
	SightsByTrip(ctx *fasthttp.RequestCtx)
}

type tripGatewayDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.TripGatewayUseCase
}

func NewTripGetewayDelivery(
	errorAdapter error_adapter.HttpAdapter,
	manager usecase.TripGatewayUseCase,
) TripGatewayDelivery {
	return &tripGatewayDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}

func (s *tripGatewayDelivery) Trip(ctx *fasthttp.RequestCtx) {
	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	trip, err := s.manager.GetTripById(ctx, param, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(trip)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *tripGatewayDelivery) AddTrip(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	trip := new(models.Trip)
	if err := json.Unmarshal(ctx.PostBody(), &trip); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	trip, err := s.manager.AddTrip(ctx, trip, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(trip)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *tripGatewayDelivery) UpdateTrip(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	trip := new(models.Trip)
	if err := json.Unmarshal(ctx.PostBody(), &trip); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	trip, err := s.manager.UpdateTrip(ctx, param, trip, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(trip)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *tripGatewayDelivery) DeleteTrip(ctx *fasthttp.RequestCtx) {

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	err := s.manager.DeleteTrip(ctx, param, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)

	response := map[string]int{"status": fasthttp.StatusOK}
	bytes, err := json.Marshal(response)
	if err != nil {
		return
	}
	ctx.Write(bytes)
}

func (s *tripGatewayDelivery) Album(ctx *fasthttp.RequestCtx) {
	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	album, err := s.manager.GetAlbumById(ctx, param, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(album)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *tripGatewayDelivery) AddAlbum(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	album := new(models.Album)
	if err := json.Unmarshal(ctx.PostBody(), &album); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	album, err := s.manager.AddAlbum(ctx, album, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(album)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *tripGatewayDelivery) UpdateAlbum(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	album := new(models.Album)
	if err := json.Unmarshal(ctx.PostBody(), &album); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	album, err := s.manager.UpdateAlbum(ctx, param, album, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(album)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *tripGatewayDelivery) DeleteAlbum(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	err := s.manager.DeleteAlbum(ctx, param, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)

	response := map[string]int{"status": fasthttp.StatusOK}
	bytes, err := json.Marshal(response)
	if err != nil {
		return
	}
	ctx.Write(bytes)
}

func (s *tripGatewayDelivery) UploadPhoto(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	formFile, err := ctx.FormFile(cnst.FormFileAlbumKey)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	filename := strconv.Itoa(param) + "_" + strconv.Itoa(userID) + formFile.Filename

	err = fasthttp.SaveMultipartFile(formFile, filename)
	if err != nil {
		fmt.Println(err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)

}

func (s *tripGatewayDelivery) SightsByTrip(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))

	ids, err := s.manager.SightsByTrip(ctx, param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(ids)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}
