package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"snakealive/m/internal/gateway/trip/usecase"
	"snakealive/m/internal/models"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
	auth_service "snakealive/m/pkg/services/auth"
	mock_auth_service "snakealive/m/pkg/services/auth/mock"
	mock_sight_service "snakealive/m/pkg/services/sight/mock"
	trip_service "snakealive/m/pkg/services/trip"
	mock_trip_service "snakealive/m/pkg/services/trip/mock"
)

type Test struct {
	Prepare func(
		trip *mock_trip_service.MockTripServiceClient,
		sight *mock_sight_service.MockSightServiceClient,
		auth *mock_auth_service.MockAuthServiceClient,
		)
	Run     func(d TripGatewayDelivery, t *testing.T)
}

const (
	defaultUserID = 1
	cookie = "cookie"

	defaultCountry = "a"
	defaultID      = 1
)

var (
	someError       = errors.New("error")
	defaultErrorMsg = "msg"
	defaultError    = error_adapter.HttpError{
		MSG:  defaultErrorMsg,
		Code: http.StatusTeapot,
	}

	tripD = trip_service.Trip{
		Id:          1,
		Title:       "1",
		Description: "1",
		Sights:      []*trip_service.Sight{
			{
				Id:          1,
				Name:        "1",
				Country:     "1",
				Rating:      1,
				Description: "1",
				Day:         1,
			},
		},
		Albums:      []*trip_service.Album{
			{
				Id:          1,
				TripId:      1,
				Author:      1,
				Title:       "1",
				Description: "1",
			},
		},
		Users:       []int64{1},
	}
	tripResp =models.TripWithUserInfo{
		Id:          1,
		Title:       "1",
		Description: "1",
		Sights:      []models.Place{
			{
				Id:          1,
				Name:        "1",
				Country:     "1",
				Rating:      1,
				Description: "1",
				Day:         1,
			},
		},
		Albums:      []models.Album{
			{
				Id:          1,
				TripId:      0,
				UserId:      0,
				Title:       "1",
				Description: "1",
			},
		},
		Users:       []models.UserInfo{
			{
				Id:      1,
				Name:    "1",
				Surname: "1",
				Avatar:  "1",
			},
		},
	}

	album = models.Album{
		Id:          1,
		TripId:      1,
		UserId:      1,
		Title:       "1",
		Description: "1",
		Photos:      []string{},
	}
)

var (
	tests = []Test{
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "asd")
				d.Trip(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().GetTrip(gomock.Any(), &trip_service.TripRequest{
					TripId: defaultID,
					UserId: defaultUserID,
				}).Return(nil, someError)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.Trip(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusNotFound)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().GetTrip(gomock.Any(), &trip_service.TripRequest{
					TripId: defaultID,
					UserId: defaultUserID,
				}).Return(&trip_service.Trip{}, nil)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.Trip(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusNotFound)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().GetTrip(gomock.Any(), &trip_service.TripRequest{
					TripId: defaultID,
					UserId: defaultUserID,
				}).Return(&tripD, nil)
				auth.EXPECT().GetUserInfo(gomock.Any(), &auth_service.GetUserRequest{Id: 1}).Return(nil, someError)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.Trip(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusNotFound)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().GetTrip(gomock.Any(), &trip_service.TripRequest{
					TripId: defaultID,
					UserId: defaultUserID,
				}).Return(&tripD, nil)
				auth.EXPECT().GetUserInfo(gomock.Any(), &auth_service.GetUserRequest{Id: 1}).Return(
					&auth_service.UserInfo{
						UserId:  1,
						Name:    "1",
						Surname: "1",
						Image:   "1",
					}, nil,
				)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.Trip(ctx)

				var trip models.TripWithUserInfo
				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
				getResp(ctx,t,  &trip)
				assert.Equal(t, trip, tripResp)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				d.AddAlbum(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().AddAlbum(gomock.Any(), gomock.Any()).Return(nil, someError)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx ,t, album)
				d.AddAlbum(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().AddAlbum(gomock.Any(), gomock.Any()).Return(&trip_service.Album{
					Id:          0,
					TripId:      0,
					Author:      0,
					Title:       "",
					Description: "",
					Photos:      []string{},
				}, nil)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				setBody(ctx ,t, album)
				d.AddAlbum(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1asd")
				d.Album(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().GetAlbum(gomock.Any(), &trip_service.AlbumRequest{
					AlbumId: 1,
					UserId:  1,
				}).Return(nil, someError)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.Album(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusNotFound)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().GetAlbum(gomock.Any(), &trip_service.AlbumRequest{
					AlbumId: 1,
					UserId:  1,
				}).Return(&trip_service.Album{
					Id:          1,
					TripId:      0,
					Author:      0,
					Title:       "",
					Description: "",
					Photos:      []string{},
				}, nil)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.Album(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},

		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().DeleteAlbum(gomock.Any(), &trip_service.AlbumRequest{
					AlbumId: 1,
					UserId:  1,
				}).Return(nil, someError)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.DeleteAlbum(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(trip *mock_trip_service.MockTripServiceClient, sight *mock_sight_service.MockSightServiceClient, auth *mock_auth_service.MockAuthServiceClient) {
				trip.EXPECT().DeleteAlbum(gomock.Any(), &trip_service.AlbumRequest{
					AlbumId: 1,
					UserId:  1,
				}).Return(nil, nil)
			},
			Run: func(d TripGatewayDelivery, t *testing.T) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.DeleteAlbum(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
	}
)

func TestDelivery(t *testing.T) {
	for i := range tests {
		d, trip, sight, auth := prepare(t)
		tests[i].Prepare(trip, sight, auth)
		tests[i].Run(d, t)
	}
}

func prepare(t *testing.T) (d TripGatewayDelivery,
	trip *mock_trip_service.MockTripServiceClient,
	sight *mock_sight_service.MockSightServiceClient,
	auth *mock_auth_service.MockAuthServiceClient,
) {
	ctrl := gomock.NewController(t)
	sight = mock_sight_service.NewMockSightServiceClient(ctrl)
	trip = mock_trip_service.NewMockTripServiceClient(ctrl)
	auth = mock_auth_service.NewMockAuthServiceClient(ctrl)

	d = NewTripGetewayDelivery(
		error_adapter.NewErrorToHttpAdapter(map[error]error_adapter.HttpError{}, defaultError),
		usecase.NewTripGatewayUseCase(trip, sight, auth),
	)

	return
}

func getCtx() *fasthttp.RequestCtx {
	ctx := &fasthttp.RequestCtx{}
	setUserDefaultCtx(ctx)
	return ctx
}

func setUserDefaultCtx(ctx *fasthttp.RequestCtx) {
	ctx.SetUserValue(cnst.UserIDContextKey, defaultUserID)
	ctx.Request.Header.SetCookie(cnst.CookieName, cookie)
}

func setBody(ctx *fasthttp.RequestCtx, t *testing.T, val interface{}) {
	b, err := json.Marshal(val)
	assert.NoError(t, err)

	ctx.Request.SetBody(b)
}

func getResp(ctx *fasthttp.RequestCtx, t *testing.T, val interface{}) {
	assert.NoError(t, json.Unmarshal(ctx.Response.Body(), val))
}
