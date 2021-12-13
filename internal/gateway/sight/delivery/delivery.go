package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"snakealive/m/internal/gateway/sight/usecase"
	"snakealive/m/internal/models"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"

	"github.com/valyala/fasthttp"
)

type SightDelivery interface {
	GetSightByID(ctx *fasthttp.RequestCtx)
	GetSightByCountry(ctx *fasthttp.RequestCtx)
	SearchSights(ctx *fasthttp.RequestCtx)
	GetSightByTag(ctx *fasthttp.RequestCtx)
	GetTags(ctx *fasthttp.RequestCtx)
	GetRandomTags(ctx *fasthttp.RequestCtx)
}

type sightDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.SightUseCase
}

func (s *sightDelivery) GetTags(ctx *fasthttp.RequestCtx) {
	response, err := s.manager.GetTags(ctx)
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

func (s *sightDelivery) GetRandomTags(ctx *fasthttp.RequestCtx) {
	response, err := s.manager.GetRandomTags(ctx)
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

	b, _ := easyjson.Marshal(response)
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

func (s *sightDelivery) GetSightByTag(ctx *fasthttp.RequestCtx) {
	tag := string(ctx.QueryArgs().Peek("tag"))
	if tag == "" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	tagID, err := strconv.Atoi(tag)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	response, err := s.manager.GetSightsByTag(ctx, tagID)
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

func (s *sightDelivery) SearchSights(ctx *fasthttp.RequestCtx) {
	var request = &models.SearchSights{}
	if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := s.manager.SearchSights(ctx, request)
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
	manager usecase.SightUseCase,
) SightDelivery {
	return &sightDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
