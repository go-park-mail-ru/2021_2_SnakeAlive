package delivery

import (
	"encoding/json"
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
	Update(ctx *fasthttp.RequestCtx)
	Delete(ctx *fasthttp.RequestCtx)
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

	userID := ctx.UserValue(cnst.UserIDContextKey).(int) //check middleware

	trip, err := s.manager.GetById(ctx, param, userID)
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

	trip, err := s.manager.Add(ctx, trip, userID)
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

func (s *tripGatewayDelivery) Update(ctx *fasthttp.RequestCtx) {

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	trip := new(models.Trip)
	if err := json.Unmarshal(ctx.PostBody(), &trip); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	trip, err := s.manager.Update(ctx, param, trip, userID)
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

func (s *tripGatewayDelivery) Delete(ctx *fasthttp.RequestCtx) {

	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	err := s.manager.Delete(ctx, param, userID)
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
