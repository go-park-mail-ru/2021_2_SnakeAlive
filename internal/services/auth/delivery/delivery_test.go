package delivery

import (
	"context"
	service_mocks "snakealive/m/internal/mocks"
	"snakealive/m/internal/services/auth/models"
	user_usecase "snakealive/m/internal/services/auth/usecase"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/grpc_errors"
	"snakealive/m/pkg/hasher"
	auth_service "snakealive/m/pkg/services/auth"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestHandler_ValidateSession(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	userID := 1
	hash := "string_hash"
	request := &auth_service.Session{
		Token:  "??",
		Cookie: hash,
	}
	expectedResponce := &auth_service.ValidateSessionResponse{
		UserId: int64(userID),
	}

	mockValidateSession := func(r *service_mocks.MockAuthRepository, ctx context.Context, id int, hash string) {
		r.EXPECT().ValidateUserSession(ctx, hash).Return(int64(id), nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	userRepo := service_mocks.NewMockAuthRepository(c)
	mockValidateSession(userRepo, ctx, userID, hash)

	userUsecase := user_usecase.NewAuthUseCase(hasher.NewHasher(5), userRepo)
	userDelivery := NewAuthDelivery(userUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, err := userDelivery.ValidateSession(ctx, request)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_Logout(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	hash := "string_hash"
	request := &auth_service.Session{
		Token:  "??",
		Cookie: hash,
	}

	mockRemoveSession := func(r *service_mocks.MockAuthRepository, ctx context.Context, hash string) {
		r.EXPECT().RemoveUserSession(ctx, hash).Return(nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	userRepo := service_mocks.NewMockAuthRepository(c)
	mockRemoveSession(userRepo, ctx, hash)

	userUsecase := user_usecase.NewAuthUseCase(hasher.NewHasher(5), userRepo)
	userDelivery := NewAuthDelivery(userUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	_, err := userDelivery.LogoutUser(ctx, request)

	assert.Nil(t, err)
}

func TestHandler_Login(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	user := &models.User{
		ID:       1,
		Email:    "email@mail.ru",
		Password: "pass",
	}
	userID := 1
	hash := "string_hash"
	request := &auth_service.LoginRequest{
		Email:    "email@mail.ru",
		Password: "pass",
	}
	expectedResponce := &auth_service.LoginResponse{
		Cookie: "",
		Token:  "",
	}

	mockGetUserByEmail := func(r *service_mocks.MockAuthRepository, ctx context.Context, user *models.User, email string) {
		r.EXPECT().GetUserByEmail(ctx, email).Return(user, nil).AnyTimes()
	}
	mockCreateUserSession := func(r *service_mocks.MockAuthRepository, ctx context.Context, id int, hash string) {
		r.EXPECT().CreateUserSession(ctx, id, hash).Return(nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	userRepo := service_mocks.NewMockAuthRepository(c)
	mockGetUserByEmail(userRepo, ctx, user, user.Email)
	mockCreateUserSession(userRepo, ctx, userID, hash)

	userUsecase := user_usecase.NewAuthUseCase(hasher.NewHasher(1), userRepo)
	userDelivery := NewAuthDelivery(userUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, _ := userDelivery.LoginUser(ctx, request)

	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_GetUser(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	user := &models.User{
		ID:       1,
		Email:    "email@mail.ru",
		Password: "pass",
	}
	request := &auth_service.GetUserRequest{
		Id: 1,
	}
	expectedResponce := &auth_service.GetUserResponse{
		Email: "email@mail.ru",
	}

	mockGetUserById := func(r *service_mocks.MockAuthRepository, ctx context.Context, user *models.User, id int64) {
		r.EXPECT().GetUserByID(ctx, id).Return(user, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	userRepo := service_mocks.NewMockAuthRepository(c)
	mockGetUserById(userRepo, ctx, user, user.ID)

	userUsecase := user_usecase.NewAuthUseCase(hasher.NewHasher(1), userRepo)
	userDelivery := NewAuthDelivery(userUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, _ := userDelivery.GetUser(ctx, request)

	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_GetUserByEmail(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	user := &models.User{
		ID:       1,
		Email:    "email@mail.ru",
		Password: "pass",
	}
	request := &auth_service.UserEmailRequest{
		Email: user.Email,
	}
	expectedResponce := &auth_service.UserId{
		Id: user.ID,
	}

	mockGetUserById := func(r *service_mocks.MockAuthRepository, ctx context.Context, user *models.User, email string) {
		r.EXPECT().GetUserByEmail(ctx, email).Return(user, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	userRepo := service_mocks.NewMockAuthRepository(c)
	mockGetUserById(userRepo, ctx, user, user.Email)

	userUsecase := user_usecase.NewAuthUseCase(hasher.NewHasher(1), userRepo)
	userDelivery := NewAuthDelivery(userUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, _ := userDelivery.GetUserByEmail(ctx, request)

	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_GetUserInfo(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	user := &models.User{
		ID:       1,
		Email:    "email@mail.ru",
		Password: "pass",
	}
	request := &auth_service.GetUserRequest{
		Id: 1,
	}
	expectedResponce := &auth_service.UserInfo{
		UserId: 1,
	}

	mockGetUserInfo := func(r *service_mocks.MockAuthRepository, ctx context.Context, user *models.User, id int) {
		r.EXPECT().GetUserInfo(ctx, id).Return(user, nil).AnyTimes()
	}

	c := gomock.NewController(t)
	defer c.Finish()

	userRepo := service_mocks.NewMockAuthRepository(c)
	mockGetUserInfo(userRepo, ctx, user, int(user.ID))

	userUsecase := user_usecase.NewAuthUseCase(hasher.NewHasher(1), userRepo)
	userDelivery := NewAuthDelivery(userUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, _ := userDelivery.GetUserInfo(ctx, request)

	assert.Equal(t, expectedResponce, responce)
}

func TestHandler_Register(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	user := &models.User{
		Email:    "email@mail.ru",
		Password: "",
	}

	newUser := &models.User{
		ID:    1,
		Email: "email@mail.ru",
	}

	request := &auth_service.RegisterRequest{
		Email:    user.Email,
		Password: "",
	}

	c := gomock.NewController(t)
	defer c.Finish()

	userRepo := service_mocks.NewMockAuthRepository(c)
	userRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(newUser, nil).AnyTimes()
	userRepo.EXPECT().CreateUserSession(ctx, newUser.ID, gomock.Any()).Return(nil).AnyTimes()

	userUsecase := user_usecase.NewAuthUseCase(hasher.NewHasher(1), userRepo)
	userDelivery := NewAuthDelivery(userUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	_, err := userDelivery.RegisterUser(ctx, request)

	assert.Nil(t, err)
}

func TestHandler_UpdateUser(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	user := &models.User{
		Email:    "email@mail.ru",
		Password: "",
	}

	request := &auth_service.UpdateUserRequest{
		Email:    user.Email,
		Password: "",
	}

	expectedResponce := &auth_service.GetUserResponse{
		Email: "email@mail.ru",
	}

	c := gomock.NewController(t)
	defer c.Finish()

	userRepo := service_mocks.NewMockAuthRepository(c)
	userRepo.EXPECT().UpdateUser(ctx, user).Return(user, nil).AnyTimes()

	userUsecase := user_usecase.NewAuthUseCase(hasher.NewHasher(1), userRepo)
	userDelivery := NewAuthDelivery(userUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	responce, err := userDelivery.UpdateUser(ctx, request)

	assert.Equal(t, expectedResponce, responce)
	assert.Nil(t, err)
}
