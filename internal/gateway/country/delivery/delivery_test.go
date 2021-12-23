package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	mock_repository "snakealive/m/internal/gateway/country/repository/mock"
	"snakealive/m/internal/gateway/country/usecase"
	"snakealive/m/internal/models"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
)

type Test struct {
	Prepare func(repo *mock_repository.MockCountryStorage)
	Run     func(t *testing.T, d CountryDelivery)
}

const (
	defaultUserID = 1
	cookie        = "cookie"
)

var (
	someError = errors.New("error")

	defaultErrorMsg = "msg"
	defaultError    = error_adapter.HttpError{
		MSG:  defaultErrorMsg,
		Code: http.StatusTeapot,
	}

	tests = []Test{
		{
			Prepare: func(repo *mock_repository.MockCountryStorage) {
				repo.EXPECT().GetCountriesList(gomock.Any()).Return(models.Countries{}, someError)
			},
			Run: func(t *testing.T, d CountryDelivery) {
				ctx := getCtx()
				d.ListCountries(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(repo *mock_repository.MockCountryStorage) {
				repo.EXPECT().GetCountriesList(gomock.Any()).Return(models.Countries{}, nil)
			},
			Run: func(t *testing.T, d CountryDelivery) {
				ctx := getCtx()
				d.ListCountries(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
		{
			Prepare: func(repo *mock_repository.MockCountryStorage) {
				repo.EXPECT().GetById(gomock.Any(), 1).Return(models.Country{}, someError)
			},
			Run: func(t *testing.T, d CountryDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.GetCountryByID(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(repo *mock_repository.MockCountryStorage) {
				repo.EXPECT().GetById(gomock.Any(), 1).Return(models.Country{}, nil)
			},
			Run: func(t *testing.T, d CountryDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.GetCountryByID(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
		{
			Prepare: func(repo *mock_repository.MockCountryStorage) {
				repo.EXPECT().GetById(gomock.Any(), 1).Return(models.Country{}, someError)
			},
			Run: func(t *testing.T, d CountryDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("name", "1")
				d.GetCountryByName(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusTeapot)
				assert.Equal(t, string(ctx.Response.Body()), defaultErrorMsg)
			},
		},
		{
			Prepare: func(repo *mock_repository.MockCountryStorage) {
				repo.EXPECT().GetById(gomock.Any(), 1).Return(models.Country{}, nil)
			},
			Run: func(t *testing.T, d CountryDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("name", "1")
				d.GetCountryByName(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
	}
)

func TestDelivery(t *testing.T) {
	for i := range tests {
		d, cli := prepare(t)
		tests[i].Prepare(cli)
		tests[i].Run(t, d)
	}
}

func prepare(t *testing.T) (d CountryDelivery, repo *mock_repository.MockCountryStorage) {
	ctrl := gomock.NewController(t)
	repo = mock_repository.NewMockCountryStorage(ctrl)
	d = NewCountryDelivery(usecase.NewCountryUsecase(repo), error_adapter.NewErrorToHttpAdapter(map[error]error_adapter.HttpError{}, defaultError))
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
