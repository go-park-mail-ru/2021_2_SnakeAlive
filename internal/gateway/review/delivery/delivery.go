package delivery

import (
	"encoding/json"
	"strconv"

	"snakealive/m/internal/gateway/review/usecase"
	"snakealive/m/internal/models"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)

type ReviewGatewayDelivery interface {
	ReviewsByPlace(ctx *fasthttp.RequestCtx)
	AddReviewToPlace(ctx *fasthttp.RequestCtx)
	DelReview(ctx *fasthttp.RequestCtx)
}

type reviewGatewayDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.ReviewGatewayUseCase
}

func NewReviewGatewayDelivery(
	errorAdapter error_adapter.HttpAdapter,
	manager usecase.ReviewGatewayUseCase,
) ReviewGatewayDelivery {
	return &reviewGatewayDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}

func (d *reviewGatewayDelivery) ReviewsByPlace(ctx *fasthttp.RequestCtx) {
	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	skip, err := strconv.Atoi(string(ctx.QueryArgs().Peek("skip")))
	if err != nil {
		skip = 0
	}

	limit, err := strconv.Atoi(string(ctx.QueryArgs().Peek("limit")))
	if err != nil {
		limit = cnst.DefaultSkip
	}

	review, err := d.manager.GetReviewsListByPlaceId(ctx, param, limit, skip)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(review)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	_, _ = ctx.Write(bytes)
}
func (d *reviewGatewayDelivery) AddReviewToPlace(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	review := new(models.Review)

	if err := easyjson.Unmarshal(ctx.PostBody(), review); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	review, err := d.manager.Add(ctx, review, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	bytes, err := easyjson.Marshal(review)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	_, _ = ctx.Write(bytes)
}
func (d *reviewGatewayDelivery) DelReview(ctx *fasthttp.RequestCtx) {
	param, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userID := ctx.UserValue(cnst.UserIDContextKey).(int)

	err = d.manager.Delete(ctx, param, userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)

}
