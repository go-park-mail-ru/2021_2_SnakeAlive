package delivery

import (
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	sight_service "snakealive/m/pkg/services/sight"
	"testing"

	service_mocks "snakealive/m/internal/mocks"

	"snakealive/m/internal/services/sight/models"
	sight_usecase "snakealive/m/internal/services/sight/usecase"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestHandler_GetSight(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	id := 1
	sight := models.Sight{
		Id:     id,
		Name:   "name",
		Tags:   []string{"tag1"},
		Photos: []string{"photo.jpg"},
	}

	request := &sight_service.GetSightRequest{
		Id: int64(id),
	}
	expectedResponce := &sight_service.GetSightResponse{
		Sight: &sight_service.Sight{
			Id:     int64(id),
			Name:   sight.Name,
			Tags:   sight.Tags,
			Photos: sight.Photos,
		},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	sightRepo := service_mocks.NewMockSightRepository(c)
	sightRepo.EXPECT().GetSightByID(ctx, id).Return(sight, nil).AnyTimes()

	sightUsecase := sight_usecase.NewSightUsecase(sightRepo)
	sightDelivery := NewSightDelivery(sightUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, err := sightDelivery.GetSight(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_GetSights(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	id := 1
	sight := models.Sight{
		Id:      id,
		Name:    "name",
		Country: "Russia",
		Tags:    []string{"tag1"},
		Photos:  []string{"photo.jpg"},
	}
	protoSight := &sight_service.Sight{
		Id:      int64(id),
		Name:    sight.Name,
		Country: sight.Country,
		Tags:    sight.Tags,
		Photos:  sight.Photos,
	}

	request := &sight_service.GetSightsRequest{
		CountryName: sight.Country,
	}
	expectedResponce := &sight_service.GetSightsReponse{
		Sights: []*sight_service.Sight{protoSight},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	sightRepo := service_mocks.NewMockSightRepository(c)
	sightRepo.EXPECT().GetSightsByCountry(ctx, sight.Country).Return([]models.Sight{sight}, nil).AnyTimes()

	sightUsecase := sight_usecase.NewSightUsecase(sightRepo)
	sightDelivery := NewSightDelivery(sightUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, err := sightDelivery.GetSights(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_GetSightsByIDs(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	id := 1
	sight := models.Sight{
		Id:      id,
		Name:    "name",
		Country: "Russia",
		Tags:    []string{"tag1"},
		Photos:  []string{"photo.jpg"},
	}
	protoSight := &sight_service.Sight{
		Id:      int64(id),
		Name:    sight.Name,
		Country: sight.Country,
		Tags:    sight.Tags,
		Photos:  sight.Photos,
	}

	request := &sight_service.GetSightsByIDsRequest{
		Ids: []int64{protoSight.Id},
	}
	expectedResponce := &sight_service.GetSightsByIDsResponse{
		Sights: []*sight_service.Sight{protoSight},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	sightRepo := service_mocks.NewMockSightRepository(c)
	sightRepo.EXPECT().GetSightByIDs(ctx, []int64{protoSight.Id}).Return([]models.Sight{sight}, nil).AnyTimes()

	sightUsecase := sight_usecase.NewSightUsecase(sightRepo)
	sightDelivery := NewSightDelivery(sightUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, err := sightDelivery.GetSightsByIDs(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_SearchSights(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	id := 1
	sight := models.Sight{
		Id:      id,
		Name:    "name",
		Country: "Russia",
		Tags:    []string{"tag1"},
		Photos:  []string{"photo.jpg"},
	}
	protoSight := &sight_service.Sight{
		Id:      int64(id),
		Name:    sight.Name,
		Country: sight.Country,
		Tags:    sight.Tags,
		Photos:  sight.Photos,
	}
	sightsSearch := models.SightsSearch{
		Skip:   0,
		Limit:  10,
		Search: "search",
	}

	request := &sight_service.SearchSightRequest{
		Search: sightsSearch.Search,
		Limit:  10,
		Skip:   0,
	}
	expectedResponce := &sight_service.SearchSightResponse{
		Sights: []*sight_service.Sight{protoSight},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	sightRepo := service_mocks.NewMockSightRepository(c)
	sightRepo.EXPECT().SearchSights(ctx, &sightsSearch).Return([]models.Sight{sight}, nil).AnyTimes()

	sightUsecase := sight_usecase.NewSightUsecase(sightRepo)
	sightDelivery := NewSightDelivery(sightUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, err := sightDelivery.SearchSights(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_GetSightsTags(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	id := 1
	sight := models.Sight{
		Id:      id,
		Name:    "name",
		Country: "Russia",
		Tags:    []string{"tag1"},
		Photos:  []string{"photo.jpg"},
	}
	protoSight := &sight_service.Sight{
		Id:      int64(id),
		Name:    sight.Name,
		Country: sight.Country,
		Tags:    sight.Tags,
		Photos:  sight.Photos,
	}

	request := &sight_service.GetSightsByTagRequest{
		Tag: 1,
	}
	expectedResponce := &sight_service.GetSightsByTagResponse{
		Sights: []*sight_service.Sight{protoSight},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	sightRepo := service_mocks.NewMockSightRepository(c)
	sightRepo.EXPECT().GetSightByTag(ctx, request.Tag).Return([]models.Sight{sight}, nil).AnyTimes()

	sightUsecase := sight_usecase.NewSightUsecase(sightRepo)
	sightDelivery := NewSightDelivery(sightUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, err := sightDelivery.GetSightsByTag(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_GetTags(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	modelTag := models.Tag{
		Id:   1,
		Name: "tag",
	}

	tag := &sight_service.Tag{
		Id:   1,
		Name: "tag",
	}

	request := &sight_service.GetTagsRequest{}
	expectedResponce := &sight_service.GetTagsResponse{
		Tags: []*sight_service.Tag{tag},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	sightRepo := service_mocks.NewMockSightRepository(c)
	sightRepo.EXPECT().GetTags(ctx).Return([]models.Tag{modelTag}, nil).AnyTimes()

	sightUsecase := sight_usecase.NewSightUsecase(sightRepo)
	sightDelivery := NewSightDelivery(sightUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, err := sightDelivery.GetTags(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_GetRandomTags(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	modelTag := models.Tag{
		Id:   1,
		Name: "tag",
	}

	tag := &sight_service.Tag{
		Id:   1,
		Name: "tag",
	}

	request := &sight_service.GetTagsRequest{}
	expectedResponce := &sight_service.GetTagsResponse{
		Tags: []*sight_service.Tag{tag},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	sightRepo := service_mocks.NewMockSightRepository(c)
	sightRepo.EXPECT().GetTags(ctx).Return([]models.Tag{modelTag}, nil).AnyTimes()

	sightUsecase := sight_usecase.NewSightUsecase(sightRepo)
	sightDelivery := NewSightDelivery(sightUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, err := sightDelivery.GetRandomTags(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponce, responce)
}
