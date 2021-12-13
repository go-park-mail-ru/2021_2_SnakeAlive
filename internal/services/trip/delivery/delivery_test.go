package delivery

import (
	"context"
	"testing"

	service_mocks "snakealive/m/internal/mocks"
	"snakealive/m/internal/services/trip/models"
	trip_usecase "snakealive/m/internal/services/trip/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	trip_service "snakealive/m/pkg/services/trip"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestHandler_GetTrip(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	userID := 1
	tripID := 1
	trip := models.Trip{
		Id:          1,
		Title:       "Best trip",
		Description: "So cool",
		Sights:      []models.Place{},
		Users:       []int{userID},
	}
	expectedTrip := &trip_service.Trip{
		Id:          1,
		Title:       "Best trip",
		Description: "So cool",
		Users:       []int64{int64(userID)},
	}

	request := &trip_service.TripRequest{
		UserId: int64(userID),
		TripId: int64(tripID),
	}

	mockGetTripAuthor := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, userID int) {
		r.EXPECT().GetTripAuthors(ctx, id).Return([]int{userID}, nil).AnyTimes()
	}

	mockGetTripById := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, trip *models.Trip) {
		r.EXPECT().GetTripById(ctx, id).Return(trip, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockGetTripAuthor(tripRepo, ctx, tripID, userID)
	mockGetTripById(tripRepo, ctx, tripID, &trip)

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	aquiredTrip, err := tripDelivery.GetTrip(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedTrip, aquiredTrip)
}

func TestHandler_AddTrip(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	userID := 1
	tripID := 1
	trip := models.Trip{
		Title:       "Best trip",
		Description: "So cool",
		Sights: []models.Place{
			{Id: 1, Name: "Heh"},
		},
	}
	tripWithId := models.Trip{
		Id:          1,
		Title:       "Best trip",
		Description: "So cool",
		Sights: []models.Place{
			{Id: 1, Name: "Heh"},
		},
		Users: []int{userID},
	}
	expectedTrip := &trip_service.Trip{
		Title:       "Best trip",
		Description: "So cool",
		Sights: []*trip_service.Sight{
			{Id: 1, Name: "Heh"},
		},
	}
	expectedTripWithId := &trip_service.Trip{
		Id:          1,
		Title:       "Best trip",
		Description: "So cool",
		Sights: []*trip_service.Sight{
			{Id: 1, Name: "Heh"},
		},
		Users: []int64{int64(userID)},
	}

	request := &trip_service.ModifyTripRequest{
		UserId: int64(userID),
		Trip:   expectedTrip,
	}

	mockGetTripById := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, trip *models.Trip, newTrip *models.Trip) {
		r.EXPECT().GetTripById(ctx, id).Return(newTrip, nil).AnyTimes()
	}
	mockAddTrip := func(r *service_mocks.MockTripRepository, ctx context.Context, trip *models.Trip, id int, userID int) {
		r.EXPECT().AddTrip(ctx, gomock.Any(), userID).Return(id, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockAddTrip(tripRepo, ctx, &trip, tripID, userID)
	mockGetTripById(tripRepo, ctx, tripID, &trip, &tripWithId)

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	aquiredTrip, err := tripDelivery.AddTrip(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedTripWithId, aquiredTrip)
}

func TestHandler_UpdateTrip(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	userID := 1
	tripID := 1
	trip := models.Trip{
		Title:       "Best trip",
		Description: "So cool",
		Sights:      []models.Place{},
	}
	tripWithId := models.Trip{
		Id:          1,
		Title:       "Best trip",
		Description: "So cool",
		Sights:      []models.Place{},
		Users:       []int{userID},
	}
	expectedTripWithId := &trip_service.Trip{
		Id:          1,
		Title:       "Best trip",
		Description: "So cool",
		Users:       []int64{int64(userID)},
	}

	request := &trip_service.ModifyTripRequest{
		UserId: int64(userID),
		Trip:   expectedTripWithId,
	}

	mockGetTripAuthor := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, userID int) {
		r.EXPECT().GetTripAuthors(ctx, id).Return([]int{userID}, nil).AnyTimes()
	}
	mockGetTripById := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, trip *models.Trip, newTrip *models.Trip) {
		r.EXPECT().GetTripById(ctx, id).Return(newTrip, nil).AnyTimes()
	}
	mockUpdateTrip := func(r *service_mocks.MockTripRepository, ctx context.Context, trip *models.Trip, id int) {
		r.EXPECT().UpdateTrip(ctx, id, gomock.Any()).Return(nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockUpdateTrip(tripRepo, ctx, &trip, tripID)
	mockGetTripById(tripRepo, ctx, tripID, &trip, &tripWithId)
	mockGetTripAuthor(tripRepo, ctx, tripID, userID)

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	aquiredTrip, err := tripDelivery.UpdateTrip(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedTripWithId, aquiredTrip)
}

func TestHandler_DeleteTrip(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	userID := 1
	tripID := 1

	request := &trip_service.TripRequest{
		UserId: int64(userID),
		TripId: int64(tripID),
	}

	mockGetTripAuthor := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, userID int) {
		r.EXPECT().GetTripAuthors(ctx, id).Return([]int{userID}, nil).AnyTimes()
	}
	mockDeleteTrip := func(r *service_mocks.MockTripRepository, ctx context.Context, id int) {
		r.EXPECT().DeleteTrip(ctx, id).Return(nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockDeleteTrip(tripRepo, ctx, tripID)
	mockGetTripAuthor(tripRepo, ctx, tripID, userID)

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	_, err := tripDelivery.DeleteTrip(ctx, request)

	assert.Nil(t, err)
}

func TestHandler_GetAlbum(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	userID := 1
	albumID := 1
	album := &models.Album{
		Id:          1,
		Title:       "Best album",
		Description: "Wow so cool",
	}
	expectedAlbum := &trip_service.Album{
		Id:          1,
		Title:       "Best album",
		Description: "Wow so cool",
	}

	request := &trip_service.AlbumRequest{
		UserId:  int64(userID),
		AlbumId: int64(albumID),
	}

	mockGetAlbumAuthor := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, userID int) {
		r.EXPECT().GetAlbumAuthor(ctx, id).Return(userID, nil).AnyTimes()
	}
	mockGetAlbum := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, album *models.Album) {
		r.EXPECT().GetAlbumById(ctx, id).Return(album, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockGetAlbum(tripRepo, ctx, albumID, album)
	mockGetAlbumAuthor(tripRepo, ctx, albumID, userID)

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	aquiredAlbum, err := tripDelivery.GetAlbum(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedAlbum, aquiredAlbum)
}

func TestHandler_AddAlbum(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	userID := 1
	albumID := 1
	album := &models.Album{
		Title:       "Best album",
		Description: "Wow so cool",
	}
	albumWithId := &models.Album{
		Id:          1,
		Title:       "Best album",
		Description: "Wow so cool",
	}
	expectedAlbum := &trip_service.Album{
		Title:       "Best album",
		Description: "Wow so cool",
	}
	expectedAlbumWithId := &trip_service.Album{
		Id:          1,
		Title:       "Best album",
		Description: "Wow so cool",
	}

	request := &trip_service.ModifyAlbumRequest{
		UserId: int64(userID),
		Album:  expectedAlbum,
	}

	mockAddAlbum := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, userId int, album *models.Album) {
		r.EXPECT().AddAlbum(ctx, album, userId).Return(id, nil).AnyTimes()
	}

	mockGetAlbum := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, album *models.Album) {
		r.EXPECT().GetAlbumById(ctx, id).Return(album, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockGetAlbum(tripRepo, ctx, albumID, albumWithId)
	mockAddAlbum(tripRepo, ctx, albumID, userID, album)

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	aquiredAlbum, err := tripDelivery.AddAlbum(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedAlbumWithId, aquiredAlbum)
}

func TestHandler_UpdateAlbum(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	userID := 1
	albumID := 1
	albumWithId := &models.Album{
		Id:          1,
		Title:       "Best album",
		Description: "Wow so cool",
	}
	expectedAlbumWithId := &trip_service.Album{
		Id:          1,
		Title:       "Best album",
		Description: "Wow so cool",
	}

	request := &trip_service.ModifyAlbumRequest{
		UserId: int64(userID),
		Album:  expectedAlbumWithId,
	}

	mockGetAlbumAuthor := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, userID int) {
		r.EXPECT().GetAlbumAuthor(ctx, id).Return(userID, nil).AnyTimes()
	}

	mockUpdateAlbum := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, userId int, album *models.Album) {
		r.EXPECT().UpdateAlbum(ctx, id, album).Return(nil).AnyTimes()
	}

	mockGetAlbum := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, album *models.Album) {
		r.EXPECT().GetAlbumById(ctx, id).Return(album, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockGetAlbum(tripRepo, ctx, albumID, albumWithId)
	mockUpdateAlbum(tripRepo, ctx, albumID, userID, albumWithId)
	mockGetAlbumAuthor(tripRepo, ctx, albumWithId.Id, userID)

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	aquiredAlbum, err := tripDelivery.UpdateAlbum(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedAlbumWithId, aquiredAlbum)
}

func TestHandler_DeleteAlbum(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	userID := 1
	albumID := 1

	request := &trip_service.AlbumRequest{
		UserId:  int64(userID),
		AlbumId: int64(albumID),
	}

	mockGetAlbumAuthor := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, userID int) {
		r.EXPECT().GetAlbumAuthor(ctx, id).Return(userID, nil).AnyTimes()
	}
	mockDeleteAlbum := func(r *service_mocks.MockTripRepository, ctx context.Context, id int) {
		r.EXPECT().DeleteAlbum(ctx, id).Return(nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockDeleteAlbum(tripRepo, ctx, albumID)
	mockGetAlbumAuthor(tripRepo, ctx, albumID, userID)

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	_, err := tripDelivery.DeleteAlbum(ctx, request)

	assert.Nil(t, err)
}

func TestHandler_SightsByTrip(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	tripID := 1
	ids := []int{1, 2, 3}
	ids64 := []int64{1, 2, 3}
	protoSights := &trip_service.Sights{
		Ids: ids64,
	}

	request := &trip_service.SightsRequest{
		TripId: int64(tripID),
	}

	mockSightsByTrip := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, sights *[]int) {
		r.EXPECT().SightsByTrip(ctx, id).Return(sights, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockSightsByTrip(tripRepo, ctx, tripID, &ids)

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	sights, err := tripDelivery.SightsByTrip(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, protoSights, sights)
}

func TestHandler_TripsByUser(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	userID := 1
	tripID := 1
	trip := models.Trip{
		Id:          1,
		Title:       "Best trip",
		Description: "So cool",
		Sights:      []models.Place{},
		Users:       []int{userID},
	}
	expectedTrip := &trip_service.Trip{

		Id:          1,
		Title:       "Best trip",
		Description: "So cool",
		Users:       []int64{int64(userID)},
	}
	expectedTrips := &trip_service.Trips{
		Trips: []*trip_service.Trip{expectedTrip},
	}

	request := &trip_service.ByUserRequest{
		UserId: int64(userID),
	}

	mockGetTripAuthor := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, userID int) {
		r.EXPECT().GetTripAuthors(ctx, id).Return([]int{userID}, nil).AnyTimes()
	}

	mockGetTripByUser := func(r *service_mocks.MockTripRepository, ctx context.Context, id int, trips *[]models.Trip) {
		r.EXPECT().GetTripsByUser(ctx, id).Return(trips, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	tripRepo := service_mocks.NewMockTripRepository(c)
	mockGetTripAuthor(tripRepo, ctx, tripID, userID)
	mockGetTripByUser(tripRepo, ctx, userID, &[]models.Trip{trip})

	tripUsecase := trip_usecase.NewTripUseCase(tripRepo)
	tripDelivery := NewTripDelivery(tripUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	aquiredTrip, err := tripDelivery.GetTripsByUser(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedTrips, aquiredTrip)
}
