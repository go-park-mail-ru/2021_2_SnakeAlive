package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"snakealive/m/internal/gateway/sight/usecase"
	"snakealive/m/pkg/error_adapter"

	"github.com/valyala/fasthttp"
)

type SightDelivery interface {
	GetSightByID(ctx *fasthttp.RequestCtx)
	GetSightByCountry(ctx *fasthttp.RequestCtx)
}

type sightDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.SightGatewayUseCase
}

func (s *sightDelivery) GetSightByID(ctx *fasthttp.RequestCtx) {
	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	response, err := s.manager.GetSightByID(ctx, param)
	if err != nil {
		httpError := s.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := json.Marshal(response)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (s *sightDelivery) GetSightByCountry(ctx *fasthttp.RequestCtx) {
	param, casted := ctx.UserValue("name").(string)
	if !casted {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	response, err := s.manager.GetSightByCountry(ctx, param)
	if err != nil {
		httpError := s.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := json.Marshal(response)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func NewSightDelivery(
	errorAdapter error_adapter.HttpAdapter,
	manager usecase.SightGatewayUseCase,
) SightDelivery {
	return &sightDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
