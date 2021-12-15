package delivery

import (
	"encoding/json"
	"log"
	"snakealive/m/internal/gateway/config"
	"snakealive/m/internal/gateway/trip/usecase"
	socket "snakealive/m/internal/models"
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
	SightsByTrip(ctx *fasthttp.RequestCtx)
	TripsByUser(ctx *fasthttp.RequestCtx)
	AlbumsByUser(ctx *fasthttp.RequestCtx)
	AddTripUser(ctx *fasthttp.RequestCtx)
	ShareLink(ctx *fasthttp.RequestCtx)
	AddUserByLink(ctx *fasthttp.RequestCtx)
	SendUpdateMessage(tripId int)
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

	responceTrip, err := s.manager.AddTrip(ctx, trip, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(responceTrip)
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

	responceTrip, err := s.manager.UpdateTrip(ctx, param, trip, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(responceTrip)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)

	s.SendUpdateMessage(responceTrip.Id)
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
	ctx.Response.SetBody(bytes)
}

func (s *tripGatewayDelivery) TripsByUser(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	trips, err := s.manager.GetTripsByUser(ctx, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(trips)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(bytes)
}

func (s *tripGatewayDelivery) AlbumsByUser(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	trips, err := s.manager.GetAlbumsByUser(ctx, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(trips)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(bytes)
}

func (s *tripGatewayDelivery) AddTripUser(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	author := ctx.UserValue(cnst.UserIDContextKey).(int)

	user := new(models.TripUser)
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err := s.manager.AddTripUser(ctx, author, param, user.Email)
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

func (s *tripGatewayDelivery) ShareLink(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	author := ctx.UserValue(cnst.UserIDContextKey).(int)

	responce, err := s.manager.ShareLink(ctx, author, param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(responce)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(bytes)
}

func (s *tripGatewayDelivery) AddUserByLink(ctx *fasthttp.RequestCtx) {
	code, _ := ctx.UserValue("code").(string)
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	author := ctx.UserValue(cnst.UserIDContextKey).(int)

	responce, err := s.manager.AddUserByLink(ctx, author, id, code)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(responce)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(bytes)
}

func (s *tripGatewayDelivery) SendUpdateMessage(tripId int) {
	var cfg config.Config
	if err := cfg.Setup(); err != nil {
		log.Fatal("failed to setup cfg: ", err)
		return
	}

	requestJSON := socket.TripRequest{
		Message: "update",
		TripId:  tripId,
	}

	bytes, err := json.Marshal(requestJSON)
	if err != nil {
		return
	}

	request := fasthttp.AcquireRequest()
	request.Header.SetMethod("POST")
	request.Header.SetContentType("application/json")
	request.SetBody(bytes)

	request.SetRequestURI(cfg.WebSocketURL)
	response := fasthttp.AcquireResponse()

	fasthttp.Do(request, response)
	fasthttp.ReleaseRequest(request)
	fasthttp.ReleaseResponse(response)
}
