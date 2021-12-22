package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"snakealive/m/internal/gateway/review/usecase"
	"snakealive/m/internal/models"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
	review_service "snakealive/m/pkg/services/review"
	mock_review_service "snakealive/m/pkg/services/review/mock"
)

type Test struct {
	Prepare func(cli *mock_review_service.MockReviewServiceClient)
	Run     func(t *testing.T, d ReviewGatewayDelivery)
}

const (
	defaultUserID = 1
	cookie        = "cookie"

	defaultUserName        = "defaultUserName"
	defaultUserSurname     = "defaultUserSurname"
	defaultUserEmail       = "defaultUserEmail"
	defaultUserImage       = "defaultUserImage"
	defaultUserDescription = "defaultUserDescription"
	pass                   = "pass"
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
			Prepare: func(cli *mock_review_service.MockReviewServiceClient) {

			},
			Run: func(t *testing.T, d ReviewGatewayDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("id", "abc")
				d.ReviewsByPlace(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(cli *mock_review_service.MockReviewServiceClient) {
				cli.EXPECT().ReviewByPlace(gomock.Any(), &review_service.ReviewRequest{
					PlaceId: 1,
					Limit:   cnst.DefaultLimit,
					Skip:    0,
				}).Return(nil, someError)
			},
			Run: func(t *testing.T, d ReviewGatewayDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.ReviewsByPlace(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusNotFound)
			},
		},
		{
			Prepare: func(cli *mock_review_service.MockReviewServiceClient) {
				cli.EXPECT().ReviewByPlace(gomock.Any(), &review_service.ReviewRequest{
					PlaceId: 1,
					Limit:   cnst.DefaultLimit,
					Skip:    0,
				}).Return(&review_service.Reviews{Reviews: []*review_service.Review{
					{
						Id:        1,
						Title:     "1",
						Text:      "1",
						Rating:    1,
						UserId:    1,
						PlaceId:   1,
						CreatedAt: "1",
					},
				}}, nil)
			},
			Run: func(t *testing.T, d ReviewGatewayDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.ReviewsByPlace(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
		{
			Prepare: func(cli *mock_review_service.MockReviewServiceClient) {
			},
			Run: func(t *testing.T, d ReviewGatewayDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("id", "abc")
				d.DelReview(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(cli *mock_review_service.MockReviewServiceClient) {
				cli.EXPECT().Delete(gomock.Any(), &review_service.DeleteReviewRequest{
					ReviewId: 1,
					UserId:   1,
				}).Return(nil, someError)
			},
			Run: func(t *testing.T, d ReviewGatewayDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.DelReview(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), fasthttp.StatusNotFound)
			},
		},
		{
			Prepare: func(cli *mock_review_service.MockReviewServiceClient) {
				cli.EXPECT().Delete(gomock.Any(), &review_service.DeleteReviewRequest{
					ReviewId: 1,
					UserId:   1,
				}).Return(nil, nil)
			},
			Run: func(t *testing.T, d ReviewGatewayDelivery) {
				ctx := getCtx()
				ctx.SetUserValue("id", "1")
				d.DelReview(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), fasthttp.StatusOK)
			},
		},
		{
			Prepare: func(cli *mock_review_service.MockReviewServiceClient) {
			},
			Run: func(t *testing.T, d ReviewGatewayDelivery) {
				ctx := getCtx()
				d.AddReviewToPlace(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusBadRequest)
			},
		},
		{
			Prepare: func(cli *mock_review_service.MockReviewServiceClient) {
				cli.EXPECT().Add(gomock.Any(), gomock.Any()).Return(nil, someError)
			},
			Run: func(t *testing.T, d ReviewGatewayDelivery) {
				ctx := getCtx()
				setBody(ctx, t, models.Review{
					Id:        1,
					Title:     "1",
					Text:      "1",
					Rating:    1,
					UserId:    1,
					PlaceId:   1,
					CreatedAt: "1",
				})
				d.AddReviewToPlace(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusNotFound)
			},
		},
		{
			Prepare: func(cli *mock_review_service.MockReviewServiceClient) {
				cli.EXPECT().Add(gomock.Any(), gomock.Any()).Return(&review_service.Review{
					Id:        1,
					Title:     "1",
					Text:      "1",
					Rating:    1,
					UserId:    1,
					PlaceId:   1,
					CreatedAt: "1",
				}, nil)
			},
			Run: func(t *testing.T, d ReviewGatewayDelivery) {
				ctx := getCtx()
				setBody(ctx, t, models.Review{
					Id:        1,
					Title:     "1",
					Text:      "1",
					Rating:    1,
					UserId:    1,
					PlaceId:   1,
					CreatedAt: "1",
				})
				d.AddReviewToPlace(ctx)

				assert.Equal(t, ctx.Response.StatusCode(), http.StatusOK)
			},
		},
	}
)

func TestDelivery(t *testing.T) {
	for i := range tests {
		d, cli := prepare(t)
		tests[i].Prepare(cli)
		tests[i].Run(t,d)
	}
}

func prepare(t *testing.T) (d ReviewGatewayDelivery, cli *mock_review_service.MockReviewServiceClient) {
	ctrl := gomock.NewController(t)
	cli = mock_review_service.NewMockReviewServiceClient(ctrl)
	d = NewReviewGatewayDelivery(
		error_adapter.NewErrorToHttpAdapter(map[error]error_adapter.HttpError{}, defaultError),
		usecase.NewReviewGatewayUseCase(cli),
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
