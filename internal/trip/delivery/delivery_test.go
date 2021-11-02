package tripDelivery

import (
	"fmt"
	cu "snakealive/m/internal/cookie/usecase"
	tu "snakealive/m/internal/trip/usecase"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/domain"
	service_mocks "snakealive/m/pkg/domain/mocks"
	"strconv"
	"testing"

	cd "snakealive/m/internal/cookie/delivery"
	logs "snakealive/m/internal/logger"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestHandler_AddTrip(t *testing.T) {
	logs.BuildLogger()

	user := domain.User{
		Id:       1,
		Name:     "Александра",
		Surname:  "Волкова",
		Email:    "testtesttests@mail.ru",
		Password: "1234567890",
	}
	inputBody := `{"title": "Best trip ever", "description": "Wow wow", "days": [[ {"id": 1}, {"id": 2} ] ]}`
	inputTrip := domain.Trip{
		Id:          0,
		Title:       "Best trip ever",
		Description: "Wow wow",
		Days:        [][]domain.Place{{domain.Place{Id: 1}, domain.Place{Id: 2}}},
	}
	addedTrip := domain.Trip{
		Id:          1,
		Title:       "Best trip ever",
		Description: "Wow wow",
		Days:        [][]domain.Place{{domain.Place{Id: 1}, domain.Place{Id: 2}}},
	}
	expectedStatusCode := fasthttp.StatusOK

	mockGetUser := func(r *service_mocks.MockCookieStorage, cookie string, user domain.User) {
		r.EXPECT().Get(cookie).Return(user, nil).AnyTimes()
	}

	mockAdd := func(r *service_mocks.MockTripStorage, trip domain.Trip, user domain.User) {
		r.EXPECT().Add(trip, user).Return(1, nil).AnyTimes()
	}

	mockGetById := func(r *service_mocks.MockTripStorage, trip domain.Trip, id int) {
		r.EXPECT().GetById(id).Return(trip, nil).AnyTimes()
	}

	ctx := &fasthttp.RequestCtx{}

	c := gomock.NewController(t)
	defer c.Finish()

	cookieRepo := service_mocks.NewMockCookieStorage(c)
	cookie := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))

	tripRepo := service_mocks.NewMockTripStorage(c)
	mockGetUser(cookieRepo, cookie, user)
	mockAdd(tripRepo, inputTrip, user)
	mockGetById(tripRepo, addedTrip, addedTrip.Id)

	ctx.Request.SetBody(nil)
	ctx.Request.AppendBody([]byte(inputBody))

	ctx.Request.Header.SetCookie(cnst.CookieName, cookie)

	cookieLayer := cd.NewCookieHandler(cu.NewCookieUseCase(cookieRepo))
	tripLayer := NewTripHandler(tu.NewTripUseCase(tripRepo), cookieLayer)
	tripLayer.AddTrip(ctx)

	assert.Equal(t, expectedStatusCode, ctx.Response.Header.StatusCode())
}

func TestHandler_Trip(t *testing.T) {
	logs.BuildLogger()

	user := domain.User{
		Id:       1,
		Name:     "Александра",
		Surname:  "Волкова",
		Email:    "testtesttests@mail.ru",
		Password: "1234567890",
	}
	trip := domain.Trip{
		Id:          1,
		Title:       "Best trip ever",
		Description: "Wow wow",
		Days:        [][]domain.Place{{domain.Place{Id: 1}, domain.Place{Id: 2}}},
	}
	expectedStatusCode := fasthttp.StatusOK

	mockGetUser := func(r *service_mocks.MockCookieStorage, cookie string, user domain.User) {
		r.EXPECT().Get(cookie).Return(user, nil).AnyTimes()
	}

	mockGetById := func(r *service_mocks.MockTripStorage, trip domain.Trip, id int) {
		r.EXPECT().GetById(id).Return(trip, nil).AnyTimes()
	}

	mockTripAuthor := func(r *service_mocks.MockTripStorage, trip domain.Trip, id int) {
		r.EXPECT().GetTripAuthor(trip.Id).Return(id).AnyTimes()
	}

	ctx := &fasthttp.RequestCtx{}

	c := gomock.NewController(t)
	defer c.Finish()

	cookieRepo := service_mocks.NewMockCookieStorage(c)
	cookie := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))

	tripRepo := service_mocks.NewMockTripStorage(c)
	mockGetUser(cookieRepo, cookie, user)
	mockGetById(tripRepo, trip, trip.Id)
	mockTripAuthor(tripRepo, trip, user.Id)

	ctx.Request.SetBody(nil)
	ctx.SetUserValue("id", strconv.Itoa(trip.Id))

	ctx.Request.Header.SetCookie(cnst.CookieName, cookie)

	cookieLayer := cd.NewCookieHandler(cu.NewCookieUseCase(cookieRepo))
	tripLayer := NewTripHandler(tu.NewTripUseCase(tripRepo), cookieLayer)
	tripLayer.Trip(ctx)

	assert.Equal(t, expectedStatusCode, ctx.Response.Header.StatusCode())
}

func TestHandler_UpdateTrip(t *testing.T) {
	logs.BuildLogger()

	user := domain.User{
		Id:       1,
		Name:     "Александра",
		Surname:  "Волкова",
		Email:    "testtesttests@mail.ru",
		Password: "1234567890",
	}
	trip := domain.Trip{
		Id:          0,
		Title:       "Worst trip ever",
		Description: "Wow wow",
		Days:        [][]domain.Place{{domain.Place{Id: 1}, domain.Place{Id: 2}}},
	}
	updatedTrip := domain.Trip{
		Id:          1,
		Title:       "Worst trip ever",
		Description: "Wow wow",
		Days:        [][]domain.Place{{domain.Place{Id: 1}, domain.Place{Id: 2}}},
	}
	inputBody := `{"title": "Worst trip ever", "description": "Wow wow", "days": [[ {"id": 1}, {"id": 2} ] ]}`
	expectedStatusCode := fasthttp.StatusOK

	mockGetUser := func(r *service_mocks.MockCookieStorage, cookie string, user domain.User) {
		r.EXPECT().Get(cookie).Return(user, nil).AnyTimes()
	}

	mockGetById := func(r *service_mocks.MockTripStorage, trip domain.Trip, id int) {
		r.EXPECT().GetById(id).Return(trip, nil).AnyTimes()
	}

	mockTripAuthor := func(r *service_mocks.MockTripStorage, trip domain.Trip, id int) {
		r.EXPECT().GetTripAuthor(trip.Id).Return(id).AnyTimes()
	}

	mockUpdate := func(r *service_mocks.MockTripStorage, trip domain.Trip, id int) {
		r.EXPECT().Update(id, trip).Return(nil).AnyTimes()
	}

	ctx := &fasthttp.RequestCtx{}

	c := gomock.NewController(t)
	defer c.Finish()

	cookieRepo := service_mocks.NewMockCookieStorage(c)
	cookie := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))

	tripRepo := service_mocks.NewMockTripStorage(c)
	mockGetUser(cookieRepo, cookie, user)
	mockGetById(tripRepo, updatedTrip, updatedTrip.Id)
	mockTripAuthor(tripRepo, updatedTrip, user.Id)
	mockUpdate(tripRepo, trip, updatedTrip.Id)

	ctx.Request.SetBody(nil)
	ctx.SetUserValue("id", strconv.Itoa(updatedTrip.Id))
	ctx.Request.AppendBody([]byte(inputBody))

	ctx.Request.Header.SetCookie(cnst.CookieName, cookie)

	cookieLayer := cd.NewCookieHandler(cu.NewCookieUseCase(cookieRepo))
	tripLayer := NewTripHandler(tu.NewTripUseCase(tripRepo), cookieLayer)
	tripLayer.Update(ctx)

	assert.Equal(t, expectedStatusCode, ctx.Response.Header.StatusCode())
}

func TestHandler_DeleteTrip(t *testing.T) {
	logs.BuildLogger()

	user := domain.User{
		Id:       1,
		Name:     "Александра",
		Surname:  "Волкова",
		Email:    "testtesttests@mail.ru",
		Password: "1234567890",
	}
	trip := domain.Trip{
		Id:          1,
		Title:       "Worst trip ever",
		Description: "Wow wow",
		Days:        [][]domain.Place{{domain.Place{Id: 1}, domain.Place{Id: 2}}},
	}
	expectedStatusCode := fasthttp.StatusOK

	mockGetUser := func(r *service_mocks.MockCookieStorage, cookie string, user domain.User) {
		r.EXPECT().Get(cookie).Return(user, nil).AnyTimes()
	}

	mockTripAuthor := func(r *service_mocks.MockTripStorage, trip domain.Trip, id int) {
		r.EXPECT().GetTripAuthor(trip.Id).Return(id).AnyTimes()
	}

	mockDelete := func(r *service_mocks.MockTripStorage, id int) {
		r.EXPECT().Delete(id).Return(nil).AnyTimes()
	}

	ctx := &fasthttp.RequestCtx{}

	c := gomock.NewController(t)
	defer c.Finish()

	cookieRepo := service_mocks.NewMockCookieStorage(c)
	cookie := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))

	tripRepo := service_mocks.NewMockTripStorage(c)
	mockGetUser(cookieRepo, cookie, user)
	mockTripAuthor(tripRepo, trip, user.Id)
	mockDelete(tripRepo, trip.Id)

	ctx.Request.SetBody(nil)
	ctx.SetUserValue("id", strconv.Itoa(trip.Id))

	ctx.Request.Header.SetCookie(cnst.CookieName, cookie)

	cookieLayer := cd.NewCookieHandler(cu.NewCookieUseCase(cookieRepo))
	tripLayer := NewTripHandler(tu.NewTripUseCase(tripRepo), cookieLayer)
	tripLayer.Delete(ctx)

	assert.Equal(t, expectedStatusCode, ctx.Response.Header.StatusCode())
}
