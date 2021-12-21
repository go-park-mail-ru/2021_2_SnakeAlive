package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"snakealive/m/internal/gateway/sight/usecase"
	"snakealive/m/internal/models"
	"snakealive/m/pkg/error_adapter"
	sight_service "snakealive/m/pkg/services/sight"
	mock_sight_service "snakealive/m/pkg/services/sight/mock"
)

type Test struct {
	Prepare func(cli *mock_sight_service.MockSightServiceClient)
	Run     func(d SightDelivery, t *testing.T)
}

const (
	defaultCountry = "a"
	defaultID = 1
)

var (
	someError       = errors.New("error")
	defaultErrorMsg = "msg"
	defaultError    = error_adapter.HttpError{
		MSG:  defaultErrorMsg,
		Code: http.StatusTeapot,
	}
	tags = []*sight_service.Tag{
		{
			Id:   2,
			Name: "2",
		},
		{
			Id:   1,
			Name: "1",
		},
	}
	respTags = []models.Tag{
		{
			Id:   2,
			Name: "2",
		},
		{
			Id:   1,
			Name: "1",
		},
	}

	sights = []*sight_service.Sight{
		{
			Id:          1,
			Name:        "1",
			Country:     "1",
			Rating:      1,
			Tags:        []string{},
			Description: "1",
			Photos:      []string{},
			Lat:         1,
			Lng:         1,
		},
		{
			Id:          2,
			Name:        "2",
			Country:     "2",
			Rating:      2,
			Tags:        []string{},
			Description: "2",
			Photos:      []string{},
			Lat:         2,
			Lng:         2,
		},
	}
	respSights = []models.Sight{
		{
			Description:   "1",
			SightMetadata: models.SightMetadata{
				Id:          1,
				Name:        "1",
				Tags:        []string{},
				Photos:      []string{},
				Country:     "1",
				Rating:      1,
				Lat:         1,
				Lng:         1,
			},
		},
		{
			Description:   "2",
			SightMetadata: models.SightMetadata{
				Id:          2,
				Name:        "2",
				Tags:        []string{},
				Photos:      []string{},
				Country:     "2",
				Rating:      2,
				Lat:         2,
				Lng:         2,
			},
		},
	}

	sight = &sight_service.Sight{
		Id:          1,
		Name:        "1",
		Country:     "1",
		Rating:      1,
		Tags:        []string{},
		Description: "1",
		Photos:      []string{},
		Lat:         1,
		Lng:         1,
	}
	respSight = models.SightMetadata{
		Id:          1,
		Name:        "1",
		Description: "1",
		Tags:        []string{},
		Photos:      []string{},
		Country:     "1",
		Rating:      1,
		Lat:         1,
		Lng:         1,
	}

	searchReq = models.SearchSights{
		Tags:      []int64{1,2,3},
		Countries: []string{},
		Skip:      1,
		Limit:     2,
		Search:    "asd",
	}
)
var (

	tests = []Test{
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetTags(gomock.Any(), &sight_service.GetTagsRequest{}).Return(nil, someError)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				d.GetTags(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetTags(gomock.Any(), &sight_service.GetTagsRequest{}).Return(&sight_service.GetTagsResponse{Tags: tags}, nil)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				d.GetTags(ctx)

				var resp []models.Tag
				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
				getResp(ctx, t, &resp)
				assert.Equal(t, resp, respTags)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetRandomTags(gomock.Any(), &sight_service.GetTagsRequest{}).Return(nil, someError)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				d.GetRandomTags(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetRandomTags(gomock.Any(), &sight_service.GetTagsRequest{}).Return(&sight_service.GetTagsResponse{Tags: tags}, nil)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				d.GetRandomTags(ctx)

				var resp []models.Tag
				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
				getResp(ctx, t, &resp)
				assert.Equal(t, resp, respTags)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				d.GetSightByCountry(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetSights(gomock.Any(), &sight_service.GetSightsRequest{CountryName: defaultCountry}).Return(nil, someError)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("name", defaultCountry)
				d.GetSightByCountry(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetSights(gomock.Any(), &sight_service.GetSightsRequest{CountryName: defaultCountry}).Return(&sight_service.GetSightsReponse{Sights: sights}, nil)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("name", defaultCountry)
				d.GetSightByCountry(ctx)

				var resp []models.Sight
				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
				getResp(ctx, t, &resp)
				assert.Equal(t, resp, respSights)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				d.GetSightByTag(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.QueryArgs().Add("tag",defaultCountry)
				d.GetSightByTag(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetSightsByTag(gomock.Any(), &sight_service.GetSightsByTagRequest{Tag: int64(defaultID)}).Return(nil, someError)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.QueryArgs().Add("tag","1")
				d.GetSightByTag(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetSightsByTag(gomock.Any(), &sight_service.GetSightsByTagRequest{Tag: int64(defaultID)}).Return(&sight_service.GetSightsByTagResponse{Sights: sights}, nil)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.QueryArgs().Add("tag","1")
				d.GetSightByTag(ctx)

				var resp []models.Sight
				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
				getResp(ctx, t, &resp)
				assert.Equal(t, resp, respSights)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id","abc")
				d.GetSightByID(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetSight(gomock.Any(), &sight_service.GetSightRequest{Id: defaultID}).Return(nil, someError)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id","1")
				d.GetSightByID(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().GetSight(gomock.Any(), &sight_service.GetSightRequest{Id: defaultID}).Return(&sight_service.GetSightResponse{Sight: sight}, nil)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id","1")
				d.GetSightByID(ctx)

				var resp models.SightMetadata
				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
				getResp(ctx, t, &resp)
				assert.Equal(t, resp, respSight)
			},
		},

		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				d.SearchSights(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().SearchSights(gomock.Any(), &sight_service.SearchSightRequest{
					Search:    searchReq.Search,
					Skip:      int64(searchReq.Skip),
					Limit:     int64(searchReq.Limit),
					Countries: searchReq.Countries,
					Tags:      searchReq.Tags,
				}).Return(nil, someError)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx, t, searchReq)
				d.SearchSights(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(cli *mock_sight_service.MockSightServiceClient) {
				cli.EXPECT().SearchSights(gomock.Any(), &sight_service.SearchSightRequest{
					Search:    searchReq.Search,
					Skip:      int64(searchReq.Skip),
					Limit:     int64(searchReq.Limit),
					Countries: searchReq.Countries,
					Tags:      searchReq.Tags,
				}).Return(&sight_service.SearchSightResponse{Sights: sights}, nil)
			},
			Run: func(d SightDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx, t, searchReq)
				d.SearchSights(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
	}
)

func TestDelivery(t *testing.T) {
	for i := range tests {
		d, cli := prepare(t)
		tests[i].Prepare(cli)
		tests[i].Run(d, t)
	}
}

func prepare(t *testing.T) (d SightDelivery, cli *mock_sight_service.MockSightServiceClient) {
	ctrl := gomock.NewController(t)
	cli = mock_sight_service.NewMockSightServiceClient(ctrl)
	d = NewSightDelivery(
		error_adapter.NewErrorToHttpAdapter(map[error]error_adapter.HttpError{}, defaultError),
		usecase.NewSightGatewayUseCase(cli),
	)

	return
}

func getCtx() *fasthttp.RequestCtx {
	ctx := &fasthttp.RequestCtx{}
	return ctx
}

func setBody(ctx *fasthttp.RequestCtx, t *testing.T, val interface{}) {
	b, err := json.Marshal(val)
	assert.NoError(t, err)

	ctx.Request.SetBody(b)
}

func getResp(ctx *fasthttp.RequestCtx, t *testing.T, val interface{}) {
	assert.NoError(t, json.Unmarshal(ctx.Response.Body(), val))
}
